package models

import (
	"fmt"
	u "preposterous/utils"

	"github.com/jinzhu/gorm"
)

type Verb struct {
	gorm.Model
	German          string
	English         string
	GrammaticalCase string
	Preposition     string
}

func (verb *Verb) Validate() (map[string]interface{}, bool) {
	if len(verb.German) < 1 {
		return u.Message(false, "German verb is required"), false
	}
	if len(verb.English) < 1 {
		return u.Message(false, "English verb is required"), false
	}
	if len(verb.GrammaticalCase) != 1 {
		return u.Message(false, "Grammatical case is required"), false
	}
	if len(verb.Preposition) < 1 {
		return u.Message(false, "German preposition is required"), false
	}
	return u.Message(false, "Validation passed"), true
}

func (verb *Verb) Create() map[string]interface{} {
	fmt.Print(verb.Validate())
	if resp, ok := verb.Validate(); !ok {
		return resp
	}

	GetDB().Create(verb)

	if verb.ID <= 0 {
		return u.Message(false, "Failed to create verb, database connection down")
	}

	response := u.Message(true, "Verb has been created")
	response["verb"] = verb
	return response
}

func GetVerb(u uint) *Verb {
	verb := &Verb{}
	GetDB().Table("verbs").Where("id = ?", u).First(verb)
	if verb.German == "" {
		return nil
	}
	return verb
}

func GetVerbs() *[]Verb {
	verb := &[]Verb{}
	GetDB().Table("verbs").Find(verb)
	return verb
}
