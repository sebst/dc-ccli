import "./App.css";

import Layout from "./layout";

import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Services from "./routes/services";
import Packages from "./routes/packages";
import Logs from "./routes/logs";
import TaskManager from "./routes/taskmanager";

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route index element={<>Home(Index)</>} />
          <Route path="/services" element={<Services />} />
          <Route path="/packages" element={<Packages />} />
          <Route path="/logs" element={<Logs />} />
          <Route path="/taskmanager" element={<TaskManager />} />
        </Routes>
      </Layout>
      {/* </> */}
    </Router>
  );
}

export default App;
