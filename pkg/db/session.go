package db

import (
	"math/rand"
	"time"
)

// Session user login cookie
type Session struct {
	Token    string `gorm:"primary_key" json:"token"`
	Username string `json:"username"`
	Expires  int64  `json:"expires"`

	CreatedAt time.Time `json:"created_at"`
}

const (
	// DefaultExpires 10 hours
	DefaultExpires = 10 * 60 * 60
)

// Create save session
func (s *Session) Create() error {
	s.CreatedAt = time.Now()
	return db.Create(s).Error
}

// Get get session by token
func (s *Session) Get() error {
	return db.Where("`token` = ?", s.Token).First(s).Error
}

// Expired token whether expired
func (s *Session) Expired() bool {
	return int64(time.Now().Sub(s.CreatedAt).Seconds()) > s.Expires
}

// Delete delete token
func (s *Session) Delete() error {
	return db.Where("`token` = ?", s.Token).Delete(&Session{}).Error
}

// SessionClean clean expired session
func SessionClean() error {
	var offset = 0
	var limit = 100
	var length = 100
	for length == limit {
		var sessions []*Session
		err := db.Offset(offset).Limit(limit).Find(&sessions).Error
		if err != nil {
			return err
		}
		for _, session := range sessions {
			if !session.Expired() {
				continue
			}
			db.Where("`token` = ?", session.Token).Delete(&Session{})
		}
		length = len(sessions)
		offset += len(sessions)
	}
	return nil
}

var (
	_sources = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	_len     = 32
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenToken generate token
func GenToken() []byte {
	var ret []byte
	var count int
	var length = len(_sources)
	for count < _len {
		idx := rand.Intn(length)
		ret = append(ret, _sources[idx])
		count++
	}
	return ret
}
