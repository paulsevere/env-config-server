import jsonFetch from "json-fetch";
import { host } from "./hostConfig";

const fetchEnv = () => dispatch => {
  console.log(dispatch);
  console.log("FERTCH ENV RUNNING");
  jsonFetch(host("all"), {
    method: "GET",
    mode: "cors",
    credentials: "omit",
    headers: { "Access-Control-Allow-Origin": "*" }
  })
    .then(res => {
      console.log(res.body);
      dispatch(res.body);
    })
    .catch(console.error);
};

export default fetchEnv;
