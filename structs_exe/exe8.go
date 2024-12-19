package main

import "fmt"

// Implement a struct `Movie` with fields `Title`, `Genre`, and `Rating` (float64).
// Write a function to filter and return a slice of movies with a rating above 8.0.

type Movie struct {
	Title  string
	Genre  string
	Rating float32
}

func filterGoodMovie(movies []Movie) []Movie {
	var best_movies []Movie
	for _, movie := range movies {
		if movie.Rating > 8.0 {
			best_movies = append(best_movies, movie)
		}
	}
	return best_movies
}

func main() {
	movies := []Movie{
		{"The Shawshank Redemption", "Drama", 9.3},
		{"The Godfather", "Crime", 9.2},
		{"The Dark Knight", "Action", 9.0},
		{"Pulp Fiction", "Crime", 8.9},
		{"Inception", "Sci-Fi", 8.8},
		{"The Matrix", "Sci-Fi", 8.7},
		{"Interstellar", "Sci-Fi", 8.6},
		{"Joker", "Thriller", 8.4},
		{"Fight Club", "Drama", 8.8},
		{"Avatar", "Action", 7.8},
		{"The Lion King", "Animation", 8.5},
		{"Frozen", "Animation", 7.4},
	}
	best_movies := filterGoodMovie(movies)

	fmt.Printf("In total, movies with rating above 8.0 is %d out of %d on the list. They are listed below: \n", len(best_movies), len(movies))
	for _, movie := range best_movies {
		fmt.Printf("Title: %s, Genre: %s, Rating: %.1f\n", movie.Title, movie.Genre, movie.Rating)
	}
}
