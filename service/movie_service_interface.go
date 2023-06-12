package service

import "github.com/Nickadimas79/movies/model"

type IMovieService interface {
	GetMovies() ([]model.Movie, error)
	GetMovie(id int) (model.Movie, error)
	CreateMovie(movie model.Movie) error
	DeleteMovie(id int) error
	DeleteAllMovie() error
	UpdateMovie(id int, movie model.Movie) error
}
