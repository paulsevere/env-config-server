import jsonFetch from "json-fetch";
import { host } from "./hostConfig";

const deleteVar = variable => dispatch => {
  jsonFetch(host(`variable?name=${variable.name}&value=${variable.value}`), {
    method: "DELETE",
    mode: "cors",
    credentials: "omit",
    headers: { "Access-Control-Allow-Origin": "*" }
  })
    .then(res => {
      console.log(res);
      return res.body;
    })
    .then(dispatch);
};

export default deleteVar;
