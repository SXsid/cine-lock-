const BASE_URL = "http://localhost:8080";
const main = async () => {
  const USER_ID = Math.random().toString(36);
  let selectedMovie = 0;
  let pollTimer;
  let movieData;
  const selectedSeatArr = {};
  const movieContanier = document.getElementById("movie-container");
  const seatContainer = document.getElementById("seat-container");

  const MovieData = async () => {
    const data = await fetch(`${BASE_URL}/movies`);
    const value = await data.json();
    return value;
  };

  const chagneStatus = async (status, key, userId) => {
    const value = key.split("");
    const row = Number(value[0]);
    const col = Number(value[1]);
    try {
      await fetch(`${BASE_URL}/seat-status`, {
        method: "PATCH",
        headers: {
          "Content-Type": "application/json",
        },

        body: JSON.stringify({
          id: selectedMovie,
          row: row,
          col: col,
          status: status,
          userId: userId,
        }),
      });
    } catch (error) {
      alert(error.message);
    }
  };

  const slectedAseat = async (key) => {
    const seat = document.getElementById(key);
    if (!seat) return;
    const value = key.split("");
    const row = Number(value[0]);
    const col = Number(value[1]);
    const currentStatus = movieData[selectedMovie].Seats[row][col].Status;
    if (currentStatus === "booked" || currentStatus === "hold") {
      alert("can't select this seat");
      return;
    }
    seat.classList.add("hold");
    chagneStatus("hold", key, USER_ID);
    renderConfirmCard(key, seat.textContent);
  };
  const selectMovie = async (key) => {
    selectedMovie = key;
    init(selectedMovie);
  };
  const renderConfirmCard = (key, name) => {
    const overlaye = document.createElement("div");
    overlaye.classList.add("overlay");

    const modal = document.createElement("div");
    modal.classList.add("modal");
    const heading = document.createElement("h3");
    heading.textContent = name;
    const confirm = document.createElement("button");
    confirm.textContent = "confirm";
    confirm.classList.add("confirm");
    confirm.addEventListener("click", () => {
      chagneStatus("booked", key, USER_ID);
      selectedSeatArr[selectedMovie].push(key);
      document.body.removeChild(overlaye);
    });
    const decline = document.createElement("button");
    decline.textContent = "Decline";
    decline.classList.add("decline");
    decline.addEventListener("click", () => {
      chagneStatus("vacant", key, USER_ID);
      document.body.removeChild(overlaye);
    });
    modal.append(heading, decline, confirm);
    overlaye.appendChild(modal);
    document.body.appendChild(overlaye);
  };
  const renderMovieContainer = (movieData, selectedMovie) => {
    movieContanier.innerHTML = "";
    for (let i = 0; i < movieData.length; i++) {
      if (!(i in selectedSeatArr)) {
        selectedSeatArr[i] = [];
      }
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
        const ID = `${row}${column}`;

        const seat = document.createElement("div");
        seat.id = ID;

        const status = seatData[row][column].Status;
        seat.addEventListener("click", () => slectedAseat(`${row}${column}`));

        if (status === "booked") seat.classList.add("booked");
        if (status === "hold") seat.classList.add("hold");
        if (status === "selected") seat.classList.add("selected");
        if (selectedSeatArr[selectedMovie].includes(ID)) {
          seat.classList.add("selected");
        }
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
      const value = json.data;

      movieData[selectedMovie].Seats = value;
      renderScreen(value);
    }, 1500);
  };
  const init = async (selectedMovie) => {
    if (!movieData) {
      const result = await MovieData();
      movieData = result.data;
    }
    renderMovieContainer(movieData, selectedMovie);
    renderScreen(movieData[selectedMovie].Seats);
    poll(selectedMovie);
  };
  init(selectedMovie);
};
document.addEventListener("DOMContentLoaded", main);
