package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
)

type Movie struct {
  Name   string
  Rating float64
  Genres []string
}

type MovieRating struct {
  Movies []Movie
  Ratings map[string]Rating
}

type Rating struct {
  Count int
  Rating float64
}

func (mr *MovieRating) init(filename string) {
  data, _ := ioutil.ReadFile(filename)
  mr.Ratings = make(map[string]Rating)
  if err := json.Unmarshal(data, mr); err != nil {
    log.Fatal(err)
  }
}

func (mr *MovieRating) calculate_rating() {
  for _, movie := range mr.Movies {
    for _, genre := range movie.Genres {
      if rating, ok := mr.Ratings[genre]; ok {
        rating.Count++
        rating.Rating += movie.Rating
      }else{
        mr.Ratings[genre] = Rating{ 1, movie.Rating }
      }
    }
  }
}

func (mr *MovieRating) print_rating() {
  var avg float64
  for genre, rating := range mr.Ratings {
    avg = rating.Rating/float64(rating.Count)
    fmt.Println("Average rating for", genre, "is", avg)
  }
}

func main() {
  var mr MovieRating
  mr.init("data.json")
  mr.calculate_rating()
  mr.print_rating()
}