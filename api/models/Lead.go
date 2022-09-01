package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Lead struct {
	ID                   uint64    `gorm:"primary_key;auto_increment" json:"id"`
	ContactTitle         string    `gorm:"size:255;not null;unique" json:"content"`
	ContactFirst         string    `gorm:"size:255;not null;" json:"content"`
	ContactMiddle        string    `gorm:"size:255;not null;" json:"content"`
	ContactLast          string    `gorm:"size:255;not null;" json:"content"`
	LeadReferalSource    string    `gorm:"size:255;not null;" json:"content"`
	DateOfInitialContact string    `gorm:"size:255;not null;" json:"content"`
	Title                string    `gorm:"size:255;not null;" json:"content"`
	Company              string    `gorm:"size:255;not null;" json:"content"`
	Industry             string    `gorm:"size:255;not null;" json:"content"`
	Address              string    `gorm:"size:255;not null;" json:"content"`
	AddressStreet        string    `gorm:"size:255;not null;" json:"content"`
	AddressState         string    `gorm:"size:255;not null;" json:"content"`
	AddressZip           string    `gorm:"size:255;not null;" json:"content"`
	Country              string    `gorm:"size:255;not null;" json:"content"`
	Phone                string    `gorm:"size:255;not null;" json:"content"`
	Email                string    `gorm:"size:255;not null;" json:"content"`
	Status               string    `gorm:"size:255;not null;" json:"content"`
	Website              string    `gorm:"size:255;not null;" json:"content"`
	LinkedInProfile      string    `gorm:"size:255;not null;" json:"content"`
	BackgroundInfo       string    `gorm:"size:255;not null;" json:"content"`
	CreatedUseID         uint32    `gorm:"not null" json:"created_user_id"`
	CreatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Lead) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Company = html.EscapeString(strings.TrimSpace(p.Company))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (l *Lead) Validate() error {

	if l.Title == "" {
		return errors.New("Required Title")
	}
	if l.Company == "" {
		return errors.New("Required Content")
	}
	if l.CreatedUseID < 1 {
		return errors.New("Required Author")
	}
	return nil
}

func (l *Lead) SaveLead(db *gorm.DB) (*Lead, error) {
	var err error
	err = db.Debug().Model(&Lead{}).Create(&l).Error
	if err != nil {
		return &Lead{}, err
	}

	return l, nil
}

func (l *Lead) FindAllLeads(db *gorm.DB) (*[]Lead, error) {
	var err error
	leads := []Lead{}
	err = db.Debug().Model(&Lead{}).Limit(100).Find(&leads).Error
	if err != nil {
		return &[]Lead{}, err
	}
	return &leads, nil
}

func (l *Lead) FindLeadByID(db *gorm.DB, pid uint64) (*Lead, error) {
	var err error
	err = db.Debug().Model(&Lead{}).Where("id = ?", pid).Take(&l).Error
	if err != nil {
		return &Lead{}, err
	}

	return l, nil
}

func (l *Lead) UpdateALead(db *gorm.DB, pid uint64) (*Lead, error) {

	var err error
	db = db.Debug().Model(&Lead{}).Where("id = ?", pid).Take(&Lead{}).UpdateColumns(
		map[string]interface{}{
			"ContactTitle":         l.ContactTitle,
			"ContactFirst":         l.ContactFirst,
			"ContactMiddle":        l.ContactMiddle,
			"ContactLast":          l.ContactLast,
			"LeadReferalSource":    l.LeadReferalSource,
			"DateOfInitialContact": l.DateOfInitialContact,
			"Title":                l.Title,
			"Company":              l.Company,
			"Industry":             l.Industry,
			"Address":              l.Address,
			"AddressStreet":        l.AddressStreet,
			"AddressState":         l.AddressState,
			"AddressZip":           l.AddressZip,
			"Country":              l.Country,
			"Phone":                l.Phone,
			"Email":                l.Email,
			"Status":               l.Status,
			"Website":              l.Website,
			"LinkedInProfile":      l.LinkedInProfile,
			"BackgroundInfo":       l.BackgroundInfo,
			"CreatedUseID":         l.CreatedUseID,
			"updated_at":           time.Now(),
		},
	)
	err = db.Debug().Model(&Lead{}).Where("id = ?", pid).Take(&l).Error
	if err != nil {
		return &Lead{}, err
	}

	return l, nil
}

func (l *Lead) DeleteALead(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Lead{}).Where("id = ? and created_user_id = ?", pid, uid).Take(&Lead{}).Delete(&Lead{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Lead not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
