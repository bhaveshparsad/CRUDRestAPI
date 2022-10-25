package database

import (
	"CRUDRestAPI/model"
)

func GetAllContacts(contacts *[]model.Contact) error {
	tx := connector.Find(contacts)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetContact(contact *model.Contact, params map[string]string) error {
	tx := connector.First(&contact, params["id"])
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func CreateContact(contact *model.Contact) error {
	tx := connector.Create(&contact)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func UpdateContact(contact *model.Contact, params map[string]string) error {
	tx := connector.First(&contact, params["id"])
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DeleteContact(contact *model.Contact, params map[string]string) error {
	tx := connector.Delete(&contact, params["id"])
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
