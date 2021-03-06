# Cloud Speech API Go samples

This directory contains [Cloud Speech API](https://cloud.google.com/speech/) Go snippets.

## Snippets

### Auto Punctuation

For more details, see the [Getting Punctuation](https://cloud.google.com/speech-to-text/docs/automatic-punctuation) tutorial in the docs.

[Go Code](auto_punctuation.go)

### Enhanced Model

For more details, see the [Using Enhanced Models](https://cloud.google.com/speech-to-text/docs/enhanced-models) tutorial in the docs.

> **Caution**: If you attempt to use an enhanced model but your Google Cloud Project does not have data logging enabled, Speech-to-Text API sends a `400` HTTP code response with the status `INVALID_ARGUMENT`. You must [enable data logging](https://cloud.google.com/speech-to-text/docs/enable-data-logging) to use the enhanced speech recognition models.

[Go Code](enhanced_model.go)
