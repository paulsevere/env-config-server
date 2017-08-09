const initialState = [
  { name: "RANDOM", value: 3000 },
  { name: "PATH", value: "/usr/local/bin" }
];

const pairs = (state = initialState, action) => {
  switch (action.type) {
    case "UPDATE_VALUE":
      return state.map(
        e => (e.name === action.payload.name ? action.payload : e)
      );
    case "DELETE_VARIABLE":
      return state.filter(e => !(e.name === action.payload.name));
    default:
      return state;
  }
};

export default pairs;
