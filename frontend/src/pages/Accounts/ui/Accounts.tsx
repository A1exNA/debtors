import { useEffect, useState } from "react"

interface Account {
	id: number
	number: string
	houseId: number
	premisesType:string
	premisesNumber: string
	ownerName: string
	ownerPhone: string
}

function Accounts() {
	const [data, setData] = useState<Account[]>([])

	useEffect(() => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/accounts")
			const data = await response.json()

			setData(data.data)
		}

		fetchData()
	}, [])

	return (
		<main className="accounts">
			<section className="accounts--counter">
				<article>Лицевых счетов - {data.length}</article>
			</section>
			<section className="accounts--list">
				{data.map((item, index) => (
					<article key={index}>
						<hr />
						<p>{item.id}</p>
						<p>{item.number},</p>
						<p>{item.houseId},</p>
						<p>{item.premisesType},</p>
						<p>{item.premisesNumber},</p>
						<p>{item.ownerName},</p>
						<p>{item.ownerPhone}</p>
						<hr />
					</article>
				))}
			</section>
		</main>
	)
}


export default Accounts