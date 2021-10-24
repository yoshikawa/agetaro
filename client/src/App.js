import { BrowserRouter, Link, Route, Switch } from "react-router-dom";
import "./App.css";
// import axios from "axios";
import LoginBox from "./components/Login.jsx";

function App() {
  // const getMessage = async () => {
  //   try {
  //     axios.defaults.headers.post["Content-Type"] =
  //       "application/json;charset=utf-8";
  //     axios.defaults.headers.post["Access-Control-Allow-Origin"] = "*";
  //     const result = await axios.get(
  //       "https://cloudrun-service-4td5gl2jwa-an.a.run.app/ping"
  //     );
  //     console.log(result);
  //   } catch (error) {
  //     console.log("error!!");
  //   }
  // };

  return (
    <div className="App">
      <BrowserRouter>
        <div>
          <nav>
            <ul>
              <li>
                <Link to="/">Home</Link>
              </li>
              <li>
                <Link to="/signup">Signup</Link>
              </li>
              <li>
                <Link to="/login">Login</Link>
              </li>
            </ul>
          </nav>

          <Switch>
            <Route path="/login">
              <LoginBox />
            </Route>
          </Switch>
        </div>
      </BrowserRouter>
    </div>
  );
}

export default App;
