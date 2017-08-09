import { version } from "inferno";
import Component from "inferno-component";
import { Provider } from "inferno-redux";
import { createStore, applyMiddleware, bindActionCreators } from "redux";

import thunk from "redux-thunk";
import reducers from "./reducers";
import "./registerServiceWorker";
import Env from "./components/Env.jsx";
import "./App.css";
import { fetchEnv } from "./components/EnvActions";
const store = createStore(reducers, applyMiddleware(thunk));

class App extends Component {
  componentWillMount() {
    console.log("coool");
    store.dispatch(fetchEnv());
  }
  render() {
    return (
      <Provider store={store}>
        <div className="App">
          <Env />
        </div>
      </Provider>
    );
  }
}

export default App;
