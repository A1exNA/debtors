import { useEffect, useState } from "react"

interface House {
	id: number
	address: string
	isServiced: boolean
}

function Houses() {
	const [data, setData] = useState<House[]>([])

	useEffect(() => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/houses")
			const data = await response.json()

			setData(data.data)
		}

		fetchData()
	}, [])

	return (
		<main className="houses">
			<section className="houses--counter">
				<article>Домов - {data.length}</article>
			</section>
			<section className="houses--list">
				{data.map((item, index) => (
					<article key={index}>
						<hr />
						<p>{item.id}</p>
						<p>{item.address},</p>
						<p>{String(item.isServiced)}</p>
						<hr />
					</article>
				))}
			</section>
		</main>
	)
}


export default Houses