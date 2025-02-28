import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import './App.css'
import Signup from './components/Signup'

function App() {

  return (
    <div>
     <div className = "text-6xl mb-5">
      welcome to parcel
      <div className = "mb-5">
        please login here:
      </div>
      <div className = "flex flex-col">
        signup here:
      </div>
     </div>
    </div>
  )
}

export default App
