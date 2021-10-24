import logo from './logo.svg';
import './App.css';
import axios from "axios";

function App() {
  const getMessage = async () => {
    try {
      const result = await axios.get('https://cloudrun-service-4td5gl2jwa-an.a.run.app/ping')
      console.log(result);
    } catch (error) {
      console.log("error!!");
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
        <button onClick={() => getMessage()}>get</button>
      </header>
    </div>
  );
}

export default App;
