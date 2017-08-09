const initialState = [
  { name: "RANDOM", value: 3000 },
  { name: "PATH", value: "/usr/local/bin" }
];

const env = (state = initialState, action) => {
  console.log(action.type);
  switch (action.type) {
    case "CURRENT_ENV":
      return action.payload;
    case "UPDATE_VALUE":
      return state.map(
        e => (e.name === action.payload.name ? action.payload : e)
      );
    case "DELETE_VARIABLE":
      return state.filter(e => !(e.name === action.payload[0].name));
    default:
      return state;
  }
};

export default env;
