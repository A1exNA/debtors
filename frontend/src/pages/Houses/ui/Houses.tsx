import { useState, useEffect } from "react"


interface House {
	id: number
	address: string
	isServiced: boolean
}

function Houses() {
	const [data, setData] = useState<House[]>([])


	const [formData, setFormData] = useState<House>({
		id: 0,
		address: "",
		isServiced: false,
	})

	const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		const { name, type, value, checked } = event.target

		setFormData(prev => (
			type == "checkbox" 
				? { ...prev, [name]: checked }
				: { ...prev, [name]: value }
		))
	}

	const closeModule = (id: string) => {
		(document.getElementById(id) as HTMLDialogElement).close()
	}


	const createHouse = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/houses", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					address: formData.address,
					isServiced: formData.isServiced,
				}),
			})

			const data = await response.json()

			if (data.status === "success") {
				setFormData({
					id: 0,
					address: "",
					isServiced: false,
				})
				closeModule("houses--create")
				readHouses()
			} else {
				console.log("Error: ", data.data)
			}
		}

		fetchData()
	}

	const readHouses = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/houses", {
				method: "GET",
				headers: {
					"Content-Type": "application/json",
				},
			})

			const data = await response.json()

			setData(data.data)
		}

		fetchData()
	}

	const updateHouse = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/houses", {
				method: "PUT",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					id: formData.id,
					address: formData.address,
					isServiced: formData.isServiced
				}),
			})

			const data = await response.json()

			if (data.status === "success") {
				setFormData({
					id: 0,
					address: "",
					isServiced: false,
				})
				closeModule("houses--update")
				readHouses()
			} else {
				console.log("Error: ", data.data)
			}
		}

		fetchData()
	}

	const deleteHouse = (id: number) => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/houses", {
				method: "DELETE",
				headers: {
					"Content-Type": "application:json",
				},
				body: JSON.stringify({
					id: id
				}),
			})

			const data = await response.json()

			if (data.status === "success") {
				readHouses()
			} else {
				console.log("Error: ", data.data)
			}
		}

		fetchData()
	}


	useEffect(() => {
		readHouses()
	}, [])


	return (
		<main className="houses">
			<button
				type="button"
				onClick={() => {(document.getElementById("houses--create") as HTMLDialogElement).showModal()}}
			>
				Добавить новый дом
			</button>

			<dialog className="houses--create" id="houses--create">
				<form id="form">
					<label>Введите адрес</label>
					<input
						type="text"
						name="address"
						value={formData.address}
						onChange={handleChange}
					/><br />
					<label>Дом еще обслуживается</label>
					<input
						type="checkbox"
						name="isServiced"
						checked={formData.isServiced}
						onChange={handleChange}
					/><br />

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); createHouse() }}
					>
						Отправить
					</button>

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); closeModule("houses--create") }}
					>
						Закрыть
					</button>
				</form>
			</dialog>


			<section className="houses--counter">
				<article>Домов - {data?.length}</article>
			</section>

			<section className="houses--list">
				{data?.map((item, index) => (
					<article key={index}>
						<hr />
						<p>{item.id}</p>
						<p>{item.address},</p>
						<p>{String(item.isServiced)}</p>
						<button
							onClick={() => {
								(document.getElementById("houses--update") as HTMLDialogElement).showModal();
								setFormData({
									id: item.id,
									address: item.address,
									isServiced: item.isServiced,
								})
							}}
						>
							Изменить
						</button>
						<button
							onClick={() => (deleteHouse(item.id))}
						>
							Удалить
						</button>
						<hr />
					</article>
				))}
			</section>

			<dialog className="houses--update" id="houses--update">
				<form>
					<label>Введите адрес</label>
					<input
						type="text"
						name="address"
						value={formData.address}
						onChange={handleChange}
					/><br />
					<label>Дом еще обслуживается</label>
					<input
						type="checkbox"
						name="isServiced"
						checked={formData.isServiced}
						onChange={handleChange}
					/><br />

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); updateHouse() }}
					>
						Отправить
					</button>
					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); closeModule("houses--update") }}
					>
						Закрыть
					</button>
				</form>
			</dialog>
		</main>
	)
}


export default Houses