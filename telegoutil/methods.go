package telegoutil

import "github.com/mymmrac/telego"

// Message creates telego.SendMessageParams with required parameters
func Message(id telego.ChatID, text string) *telego.SendMessageParams {
	return &telego.SendMessageParams{
		ChatID: id,
		Text:   text,
	}
}

// Photo creates telego.SendPhotoParams with required parameters
func Photo(id telego.ChatID, photo telego.InputFile) *telego.SendPhotoParams {
	return &telego.SendPhotoParams{
		ChatID: id,
		Photo:  photo,
	}
}

// Audio creates telego.SendAudioParams with required parameters
func Audio(id telego.ChatID, audio telego.InputFile) *telego.SendAudioParams {
	return &telego.SendAudioParams{
		ChatID: id,
		Audio:  audio,
	}
}

// Document creates telego.SendDocumentParams with required parameters
func Document(id telego.ChatID, document telego.InputFile) *telego.SendDocumentParams {
	return &telego.SendDocumentParams{
		ChatID:   id,
		Document: document,
	}
}

// Video creates telego.SendVideoParams with required parameters
func Video(id telego.ChatID, video telego.InputFile) *telego.SendVideoParams {
	return &telego.SendVideoParams{
		ChatID: id,
		Video:  video,
	}
}

// Animation creates telego.SendAnimationParams with required parameters
func Animation(id telego.ChatID, animation telego.InputFile) *telego.SendAnimationParams {
	return &telego.SendAnimationParams{
		ChatID:    id,
		Animation: animation,
	}
}

// Voice creates telego.SendVoiceParams with required parameters
func Voice(id telego.ChatID, voice telego.InputFile) *telego.SendVoiceParams {
	return &telego.SendVoiceParams{
		ChatID: id,
		Voice:  voice,
	}
}

// VideoNote creates telego.SendVideoNoteParams with required parameters
func VideoNote(id telego.ChatID, videoNote telego.InputFile) *telego.SendVideoNoteParams {
	return &telego.SendVideoNoteParams{
		ChatID:    id,
		VideoNote: videoNote,
	}
}

// MediaGroup creates telego.SendMediaGroupParams with required parameters
func MediaGroup(id telego.ChatID, mediaGroups []telego.InputMedia) *telego.SendMediaGroupParams {
	return &telego.SendMediaGroupParams{
		ChatID: id,
		Media:  mediaGroups,
	}
}

// Location creates telego.SendLocationParams with required parameters
func Location(id telego.ChatID, latitude, longitude float64) *telego.SendLocationParams {
	return &telego.SendLocationParams{
		ChatID:    id,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// Venue creates telego.SendVenueParams with required parameters
func Venue(id telego.ChatID, latitude, longitude float64, title, address string) *telego.SendVenueParams {
	return &telego.SendVenueParams{
		ChatID:    id,
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// Contact creates telego.SendContactParams with required parameters
func Contact(id telego.ChatID, phoneNumber, firstName string) *telego.SendContactParams {
	return &telego.SendContactParams{
		ChatID:      id,
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// Poll creates telego.SendPollParams with required parameters
func Poll(id telego.ChatID, question string, options []string) *telego.SendPollParams {
	return &telego.SendPollParams{
		ChatID:   id,
		Question: question,
		Options:  options,
	}
}

// Dice creates telego.SendDiceParams with required parameters
// Note: Emoji isn't required, but most likely you would what to specify it
func Dice(id telego.ChatID, emoji string) *telego.SendDiceParams {
	return &telego.SendDiceParams{
		ChatID: id,
		Emoji:  emoji,
	}
}

// ChatAction creates telego.SendChatActionParams with required parameters
func ChatAction(id telego.ChatID, action string) *telego.SendChatActionParams {
	return &telego.SendChatActionParams{
		ChatID: id,
		Action: action,
	}
}

// Sticker creates telego.SendStickerParams with required parameters
func Sticker(id telego.ChatID, sticker telego.InputFile) *telego.SendStickerParams {
	return &telego.SendStickerParams{
		ChatID:  id,
		Sticker: sticker,
	}
}

// Invoice creates telego.SendInvoiceParams with required parameters
func Invoice(id telego.ChatID, title, description, payload, providerToken, currency string,
	prices []telego.LabeledPrice) *telego.SendInvoiceParams {
	return &telego.SendInvoiceParams{
		ChatID:        id,
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}
}

// Game creates telego.SendGameParams with required parameters
func Game(id int64, gameShortName string) *telego.SendGameParams {
	return &telego.SendGameParams{
		ChatID:        id,
		GameShortName: gameShortName,
	}
}