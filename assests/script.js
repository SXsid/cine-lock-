const BASE_URL = "http://localhost:8080";
const main = async () => {
  let selectedMovie = 0;
  let pollTimer;
  let movieData;
  const movieContanier = document.getElementById("movie-container");
  const seatContainer = document.getElementById("seat-container");

  const MovieData = async () => {
    const data = await fetch(`${BASE_URL}/movies`);
    const value = await data.json();
    return value;
  };

  const chagneStatus = async (status, key) => {
    await fetch(
      `${BASE_URL}/seat-status?id=${selectedMovie}&status=${status}&key=${key}`,
      {
        method: "PATCH",
      },
    );
  };

  const slectedAseat = async (key) => {
    const seat = document.getElementById(key);
    if (!seat) return;
    const value = key.split("");
    const row = Number(value[0]);
    const col = Number(value[1]);
    const currentStatus = movieData[selectedMovie].Seats[row][col].Status;
    if (currentStatus === "hold") {
      alert("alredy selected");
      return;
    }
    movieData[selectedMovie].Seats[row][col].Status = "hold";
    seat.classList.add("hold");
    chagneStatus("hold", key);
  };
  const selectMovie = async (key) => {
    selectedMovie = key;
    init(selectedMovie);
  };
  const renderMovieContainer = (movieData, selectedMovie) => {
    movieContanier.innerHTML = "";
    for (let i = 0; i < movieData.length; i++) {
      const movie = document.createElement("div");
      movie.id = i;
      movie.classList.add("movie");
      if (i === selectedMovie) {
        movie.classList.add("active");
      }
      movie.innerText = movieData[i].Name;
      movie.addEventListener("click", () => selectMovie(i));
      movieContanier.append(movie);
    }
  };
  const renderScreen = (seatData) => {
    seatContainer.innerHTML = "";
    for (let row = 0; row < seatData.length; row++) {
      const rowContainer = document.createElement("div");
      rowContainer.classList.add("row");
      for (let column = 0; column < seatData[0].length; column++) {
        const seat = document.createElement("div");
        seat.id = `${row}${column}`;
        const status = seatData[row][column].Status;
        seat.addEventListener("click", () => slectedAseat(`${row}${column}`));

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
  //polling for status
  const poll = (selectedMovie) => {
    if (pollTimer) {
      clearInterval(pollTimer);
    }
    pollTimer = setInterval(async () => {
      const data = await fetch(`${BASE_URL}/poll-status?id=${selectedMovie}`);
      const json = await data.json();
      renderScreen(json.data);
    }, 1000);
  };
  const init = async (selectedMovie) => {
    if (!movieData) {
      result = await MovieData();
      movieData = result.data;
    }
    renderMovieContainer(movieData, selectedMovie);
    renderScreen(movieData[selectedMovie].Seats);
    poll(selectedMovie);
  };
  init(selectedMovie);
};
document.addEventListener("DOMContentLoaded", main);
