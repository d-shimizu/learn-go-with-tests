package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league   League
}

// NewFileSystemPlayerStore は、ファイルシステムを使ったプレイヤーストアを生成する
func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)

	return &FileSystemPlayerStore{
		database: &tape{database},
		league:   league,
	}
}

// GetLeague は、リーグ情報を取得する
func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {

	player := f.league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	//f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}
