const BASE_URL = "http://localhost:8080";
const main = async () => {
  const movieContanier = document.getElementById("movie-container");
  const seatContainer = document.getElementById("seat-container");

  const MovieData = async () => {
    const data = await fetch(`${BASE_URL}/movies`);
    const value = await data.json();
    return value;
  };
  const renderMovieContainer = (movieData, selectedMovie) => {
    for (i = 0; i < movieData.length; i++) {
      const movie = document.createElement("div");
      movie.classList.add("movie");
      if (i === selectedMovie) {
        movie.classList.add("active");
      }
      movie.innerText = movieData[i].Name;
      movieContanier.append(movie);
    }
  };
  const renderScreen = (seatData) => {
    console.log(seatData);
    for (let row = 0; row < seatData.length; row++) {
      const rowContainer = document.createElement("div");
      rowContainer.classList.add("row");
      for (let column = 0; column < seatData[0].length; column++) {
        const seat = document.createElement("div");
        const status = seatData[row][column].Status;

        if (status === "booked") seat.classList.add("booked");
        if (status === "hold") seat.classList.add("hold");
        if (status === "selected") seat.classList.add("selected");
        seat.classList.add("seat");
        seat.innerText = seatData[row][column].Name;
        rowContainer.append(seat);
      }
      seatContainer.append(rowContainer);
    }
  };
  const init = async () => {
    movieData = await MovieData();
    renderMovieContainer(movieData.data, 0);
    renderScreen(movieData.data[0].Seats);
  };
  init();
};
document.addEventListener("DOMContentLoaded", main);
