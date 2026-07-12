import { BrowserRouter, Route, Routes, Link } from "react-router-dom"
import Houses from "../pages/Houses/ui/Houses"
import Accounts from "../pages/Accounts/ui/Accounts"
import Debts from "../pages/Debts/ui/Debts"


function App() {
  return (
    <BrowserRouter>
			<nav>
				<Link to="/">Главная</Link> | <Link to="/houses">Дома</Link> | <Link to="/accounts">Лицевые счета</Link> | <Link to="/debts">Задолженности</Link>
			</nav>
			<Routes>
				<Route path="/" element={<></>}/>
				<Route path="/houses" element={<Houses />}/>
				<Route path="/accounts" element={<Accounts />}/>
				<Route path="/debts" element={<Debts />}/>
			</Routes>
    </BrowserRouter>
  )
}

export default App
