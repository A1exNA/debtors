import { useEffect, useState } from "react"

interface Debt {
	id: number
	accountNumber: string
	reportDate: string
	openingBalance: number
	accrued: number
	paid: number
	closingBalance: number
}

function Debts() {
	const [data, setData] = useState<Debt[]>([])

	useEffect(() => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/debts")
			const data = await response.json()

			setData(data.data)
		}

		fetchData()
	}, [])

	return (
		<main className="debts">
			<section className="debts--counter">
				<article>Строк - {data.length}</article>
			</section>
			<section className="debts--list">
				{data.map((item, index) => (
					<article key={index}>
						<hr />
						<p>{item.id}</p>
						<p>{item.accountNumber},</p>
						<p>{item.reportDate},</p>
						<p>{item.openingBalance},</p>
						<p>{item.accrued},</p>
						<p>{item.paid},</p>
						<p>{item.closingBalance}</p>
						<hr />
					</article>
				))}
			</section>
		</main>
	)
}


export default Debts