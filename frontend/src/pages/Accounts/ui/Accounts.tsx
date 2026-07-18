import { useState, useEffect } from "react"


interface Account {
	id: number
	number: number
	houseId: number
	premisesType: string
	premisesNumber: string
	ownerName: string
	ownerPhone: string
}

function Accounts() {
	const [data, setData] = useState<Account[]>([])


	const [formData, setFormData] = useState<Account>({
		id: 0,
		number: 0,
		houseId: 0,
		premisesType: "",
		premisesNumber: "",
		ownerName: "",
		ownerPhone: "",
	})


	const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = event.target

		setFormData(prev => ({ ...prev, [name]: value }))
	}

	const closeModule = (id: string) => {
		(document.getElementById(id) as HTMLDialogElement).close()
	}


	const createAccount = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/accounts", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					number: Number(formData.number),
					houseId: Number(formData.houseId),
					premisesType: formData.premisesType === "" ? null : formData.premisesType,
					premisesNumber: formData.premisesNumber === "" ? null : formData.premisesNumber,
					ownerName: formData.ownerName === "" ? null : formData.ownerName,
					ownerPhone: formData.ownerPhone === "" ? null : formData.ownerPhone,
				}),
			})

			const data = await response.json()

			if (data.status === "success") {
				setFormData({
					id: 0,
					number: 0,
					houseId: 0,
					premisesType: "",
					premisesNumber: "",
					ownerName: "",
					ownerPhone: "",
				})
				closeModule("accounts--create")
				readAccounts()
			} else {
				console.error("Error: ", data.data)
			}
		}

		fetchData()
	}

	const readAccounts = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/accounts", {
				method: "GET",
				headers: {
					"Content-Type": "application/json",
				},
			})

			const data = await response.json()

			if (data.status === "success") {
				setData(data.data)
			} else {
				console.error("Error: ", data.data)
			}
		}

		fetchData()
	}

	const updateAccount = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/accounts", {
				method: "PUT",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					id: formData.id,
					number: Number(formData.number),
					houseId: Number(formData.houseId),
					premisesType: formData.premisesType === "" ? null : formData.premisesType,
					premisesNumber: formData.premisesNumber === "" ? null : formData.premisesNumber,
					ownerName: formData.ownerName === "" ? null : formData.ownerName,
					ownerPhone: formData.ownerPhone === "" ? null : formData.ownerPhone,
				}),
			})

			const data = await response.json()

			if (data.status === "success") {
				setFormData({
					id: 0,
					number: 0,
					houseId: 0,
					premisesType: "",
					premisesNumber: "",
					ownerName: "",
					ownerPhone: "",
				})
				closeModule("accounts--update")
				readAccounts()
			} else {
				console.error("Error: ", data.data)
			}
		}

		fetchData()
	}

	const deleteAccount = (id: number) => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/accounts", {
				method: "DELETE",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					id: id,
				}),
			})

			const data = await response.json()

			if (data.status === "success") {
				readAccounts()
			} else {
				console.error("Error: ", data.data)
			}
		}

		fetchData()
	}

	useEffect(() => {
		readAccounts()
	}, [])


	return (
		<main className="account">
			<button
				type="button"
				onClick={() => (document.getElementById("accounts--create") as HTMLDialogElement).showModal()}
			>
				Добавить новый лицевой счет
			</button>

			<dialog className="accounts--create" id="accounts--create">
				<form>
					<label>Введите номер лицевого счета</label>
					<input
						type="number"
						name="number"
						value={formData.number}
						onChange={handleChange}
					/><br />
					<label>Введите id дома</label>
					<input
						type="number"
						name="houseId"
						value={formData.houseId}
						onChange={handleChange}
					/><br />
					<label>Введите тип помещения</label>
					<input
						type="text"
						name="premisesType"
						value={formData.premisesType}
						onChange={handleChange}
					/><br />
					<label>Введите номер помещения</label>
					<input
						type="text"
						name="premisesNumber"
						value={formData.premisesNumber}
						onChange={handleChange}
					/><br />
					<label>Введите ФИО собственника</label>
					<input
						type="text"
						name="ownerName"
						value={formData.ownerName}
						onChange={handleChange}
					/><br />
					<label>Введите номер телефона собственника</label>
					<input
						type="text"
						name="ownerPhone"
						value={formData.ownerPhone}
						onChange={handleChange}
					/><br />

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); createAccount() }}
					>
						Отправить
					</button>

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); closeModule("accounts--create") }}
					>
						Закрыть
					</button>
				</form>
			</dialog>


			<section className="accounts--counter">
				<article>Лицевых счетов - {data?.length}</article>
			</section>

			<section className="accounts--list">
				{data?.map((item, index) => (
					<article key={index}>
						<hr />
						<p>{item.id}</p>
						<p>{item.number}</p>
						<p>{item.houseId}</p>
						<p>{item.premisesType}</p>
						<p>{item.premisesNumber}</p>
						<p>{item.ownerName}</p>
						<p>{item.ownerPhone}</p>

						<button
							type="button"
							onClick={() => {
								(document.getElementById("accounts--update") as HTMLDialogElement).showModal()
								setFormData({
									id: item.id,
									number: item.number,
									houseId: item.houseId,
									premisesType: item.premisesType ?? "",
									premisesNumber: item.premisesNumber ?? "",
									ownerName: item.ownerName ?? "",
									ownerPhone: item.ownerPhone ?? "",
								})
							}}
						>
							Изменить
						</button>

						<button
							type="button"
							onClick={() => deleteAccount(item.id)}
						>
							Удалить
						</button>
						<hr />
					</article>
				))}
			</section>

			<dialog className="accounts--update" id="accounts--update">
				<form>
					<label>Введите номер лицевого счета</label>
					<input
						type="number"
						name="number"
						value={formData.number}
						onChange={handleChange}
					/><br />
					<label>Введите id дома</label>
					<input
						type="number"
						name="houseId"
						value={formData.houseId}
						onChange={handleChange}
					/><br />
					<label>Введите тип помещения</label>
					<input
						type="text"
						name="premisesType"
						value={formData.premisesType}
						onChange={handleChange}
					/><br />
					<label>Введите номер помещения</label>
					<input
						type="text"
						name="premisesNumber"
						value={formData.premisesNumber}
						onChange={handleChange}
					/><br />
					<label>Введите ФИО собственника</label>
					<input
						type="text"
						name="ownerName"
						value={formData.ownerName}
						onChange={handleChange}
					/><br />
					<label>Введите номер телефона собственника</label>
					<input
						type="text"
						name="ownerPhone"
						value={formData.ownerPhone}
						onChange={handleChange}
					/><br />

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); updateAccount() }}
					>
						Отправить
					</button>
					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); closeModule("accounts--update") }}
					>
						Закрыть
					</button>
				</form>
			</dialog>
		</main>
	)
}


export default Accounts