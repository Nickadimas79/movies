package repository

import "github.com/Nickadimas79/movies/model"

type IMovieRepository interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) error
	DeleteMovie(id int) error
	DeleteAllMovies() error
	UpdateMovie(id int, movie model.Movie) error
}
