package model

import (
	tf "github.com/galeone/tensorflow/tensorflow/go"
)

// Embedding Defaults
const (
	EmbeddingModelTag = "bert-pretrained"
	EmbeddingOp       = "embedding"
)

// NewEmbeddings returns a pre-trained model for text embeddings
func NewEmbeddings(path string, opts ...BertOption) (Bert, error) {
	vocabPath := (path + "/" + DefaultVocabFile) // TODO os.Join
	m, err := tf.LoadSavedModel(path, []string{EmbeddingModelTag}, nil)
	if err != nil {
		return Bert{}, err
	}
	return NewBert(m, vocabPath, opts...)
}
