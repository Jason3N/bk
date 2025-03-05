import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import './App.css'
import Signup from './components/Signup'
import HomePage from "./components/Homepage"

function App() {

  return (
    <Router>
      <Routes>
        <Route path ="/" element = {<HomePage/>}/>
        <Route path ="/signup" element = {<Signup/>}/>
      </Routes>
    </Router>
  )
}

export default App
