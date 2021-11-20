package telego

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/valyala/fasthttp"
)

func TestBot_SetUpdateInterval(t *testing.T) {
	bot := &Bot{}
	ui := time.Second

	bot.SetUpdateInterval(ui)
	assert.Equal(t, ui, bot.updateInterval)
}

func TestBot_StopGettingUpdates(t *testing.T) {
	bot := &Bot{}

	bot.stopChannel = make(chan struct{})
	assert.NotPanics(t, func() {
		bot.StopLongPulling()
	})
}

func TestBot_GetUpdatesChan(t *testing.T) {
	ctrl := gomock.NewController(t)

	t.Run("success", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(data, nil).MinTimes(1)

		expectedUpdates := []Update{
			{UpdateID: 1},
			{UpdateID: 2},
		}
		setResult(t, expectedUpdates)
		m.MockAPICaller.EXPECT().
			Call(gomock.Any(), gomock.Any()).
			Return(resp, nil).MinTimes(1)

		assert.NotPanics(t, func() {
			_, err := m.Bot.GetUpdatesViaLongPulling(nil)
			assert.NoError(t, err)
			time.Sleep(time.Millisecond * 10)
			m.Bot.StopLongPulling()
			time.Sleep(time.Millisecond * 500)
		})
	})

	t.Run("error", func(t *testing.T) {
		m := newMockedBot(ctrl)

		m.MockRequestConstructor.EXPECT().
			JSONRequest(gomock.Any()).
			Return(nil, errTest).MinTimes(1)

		assert.NotPanics(t, func() {
			_, err := m.Bot.GetUpdatesViaLongPulling(nil)
			assert.NoError(t, err)
			time.Sleep(time.Millisecond * 10)
			m.Bot.StopLongPulling()
		})
	})
}

func TestBot_StartListeningForWebhook(t *testing.T) {
	b, err := NewBot(token, DefaultLogger(false, false))
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		b.StartListeningForWebhook("127.0.0.1:3000")
		time.Sleep(time.Millisecond * 10)
	})

	assert.NotPanics(t, func() {
		b.StartListeningForWebhook("test")
	})
}

func TestBot_StartListeningForWebhookTLSEmbed(t *testing.T) {
	b, err := NewBot(token, DefaultLogger(false, false))
	require.NoError(t, err)

	c, k, err := fasthttp.GenerateTestCertificate("127.0.0.1")
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		b.StartListeningForWebhookTLSEmbed("127.0.0.1:3000", c, k)
		time.Sleep(time.Millisecond * 10)
	})

	assert.NotPanics(t, func() {
		b.StartListeningForWebhookTLSEmbed("test", nil, nil)
	})
}

func TestBot_StartListeningForWebhookTLS(t *testing.T) {
	b, err := NewBot(token, DefaultLogger(false, false))
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		b.StartListeningForWebhookTLS("127.0.0.1:3000", "", "")
		time.Sleep(time.Millisecond * 10)
	})
}

func TestBot_respondWithError(t *testing.T) {
	b, err := NewBot(token, DefaultLogger(false, false))
	require.NoError(t, err)

	ctx := &fasthttp.RequestCtx{}

	b.respondWithError(ctx, errTest)
	assert.Equal(t, fasthttp.StatusBadRequest, ctx.Response.StatusCode())
}

func TestBot_StopWebhook(t *testing.T) {
	b, err := NewBot(token, DefaultLogger(false, false))
	require.NoError(t, err)

	b.stopChannel = make(chan struct{})
	assert.NotPanics(t, func() {
		err := b.StopWebhook()
		assert.NoError(t, err)
	})
}

func TestBot_GetUpdatesViaWebhook(t *testing.T) {
	b, err := NewBot(token, DefaultLogger(false, false))
	require.NoError(t, err)

	_, err = b.GetUpdatesViaWebhook("/bot")
	require.NoError(t, err)

	assert.NotPanics(t, func() {
		t.Run("invalid_path_error", func(t *testing.T) {
			ctx := &fasthttp.RequestCtx{}
			b.server.Handler(ctx)
		})

		t.Run("invalid_method_error", func(t *testing.T) {
			ctx := &fasthttp.RequestCtx{}
			ctx.Request.SetRequestURI("/bot")
			b.server.Handler(ctx)
		})
	})
}