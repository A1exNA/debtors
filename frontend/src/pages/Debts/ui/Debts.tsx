import React, { useState, useEffect } from "react"

interface Debt {
	id: number
	accountNumber: number
	reportDate: string
	openingBalance: number
	accrued: number
	paid: number
	closingBalance: number
	uploadDate: string
}

function Debts() {
	const [data, setData] = useState<Debt[]>([])

	const [formData, setFormData] = useState<Debt>({
		id: 0,
		accountNumber: 0,
		reportDate: "",
		openingBalance: 0,
		accrued: 0,
		paid: 0,
		closingBalance: 0,
		uploadDate: "",
	})

	const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = event.target

		setFormData(prev => ({ ...prev, [name]: value }))
	}

	const closeModule = (id: string) => {
		(document.getElementById(id) as HTMLDialogElement).close()
	}


	const createDebt = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/debts", {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					accountNumber: Number(formData.accountNumber),
					reportDate: new Date(formData.reportDate).toISOString(),
					openingBalance: Number(formData.openingBalance),
					accrued: Number(formData.accrued),
					paid: Number(formData.paid),
					closingBalance: Number(formData.closingBalance),
				}),
			})

			const data = await response.json()

			if (data.status === "success") {
				setFormData({
					id: 0,
					accountNumber: 0,
					reportDate: "",
					openingBalance: 0,
					accrued: 0,
					paid: 0,
					closingBalance: 0,
					uploadDate: "",
				})
				closeModule("debts--create")
				readDebts()
			}  else {
				console.error("Error: ", data.data)
			}
		}

		fetchData()
	}

	const readDebts = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/debts", {
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

	const updateDebt = () => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/debts", {
				method: "PUT",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					id: Number(formData.id),
					accountNumber: Number(formData.accountNumber),
					reportDate: new Date(formData.reportDate).toISOString(),
					openingBalance: Number(formData.openingBalance),
					accrued: Number(formData.accrued),
					paid: Number(formData.paid),
					closingBalance: Number(formData.closingBalance),
				}),
			})
			const data = await response.json()

			if (data.status === "success") {
				setFormData({
					id: 0,
					accountNumber: 0,
					reportDate: "",
					openingBalance: 0,
					accrued: 0,
					paid: 0,
					closingBalance: 0,
					uploadDate: "",
				})
				closeModule("debts--update")
				readDebts()
			} else {
				console.error("Error: ", data.data)
			}
		}

		fetchData()
	}

	const deleteDebt = (id: number) => {
		const fetchData = async () => {
			const response = await fetch("http://localhost:8080/debts", {
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
				readDebts()
			} else {
				console.error("Error: ", data.data)
			}
		}

		fetchData()
	}

	useEffect(() => {
		readDebts()
	}, [])


	return (
		<main className="debts">
			<button
				type="button"
				onClick={() => (document.getElementById("debts--create") as HTMLDialogElement).showModal()}
			>
				Добавить новую запись
			</button>

			<dialog className="debts--create" id="debts--create">
				<form id="form">
					<label>Введите номер лицевого счета</label>
					<input
						type="number"
						name="accountNumber"
						value={formData.accountNumber}
						onChange={handleChange}
					/><br />
					<label>Введите дату отчета</label>
					<input
						type="date"
						name="reportDate"
						value={formData.reportDate}
						onChange={handleChange}
					/><br />
					<label>Введите Начальное сальдо</label>
					<input
						type="number"
						name="openingBalance"
						value={formData.openingBalance}
						onChange={handleChange}
					/><br />
					<label>Введите Начисления сальдо</label>
					<input
						type="number"
						name="accrued"
						value={formData.accrued}
						onChange={handleChange}
					/><br />
					<label>Введите Платеж сальдо</label>
					<input
						type="number"
						name="paid"
						value={formData.paid}
						onChange={handleChange}
					/><br />
					<label>Введите Конечное сальдо сальдо</label>
					<input
						type="number"
						name="closingBalance"
						value={formData.closingBalance}
						onChange={handleChange}
					/><br />

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); createDebt() }}
					>
						Отправить
					</button>

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); closeModule("debts--create") }}
					>
						Закрыть
					</button>
				</form>
			</dialog>

			<section className="debts--counter">
				<article>Строк - {data.length}</article>
			</section>
			<section className="debts--list">
				{data.map((item, index) => (
					<article key={index}>
						<hr />
						<p>{item.id}</p>
						<p>{item.accountNumber}</p>
						<p>{item.reportDate.slice(0, item.reportDate.indexOf("T")).replaceAll("-", ".")}</p>
						<p>{item.openingBalance}</p>
						<p>{item.accrued}</p>
						<p>{item.paid}</p>
						<p>{item.closingBalance}</p>
						<p>{item.uploadDate}</p>

						<button
							type="button"
							onClick={() => {
								(document.getElementById("debts--update") as HTMLDialogElement).showModal()
								setFormData({
									id: item.id,
									accountNumber: item.accountNumber,
									reportDate: item.reportDate.slice(0, item.reportDate.indexOf("T")),
									openingBalance: item.openingBalance,
									accrued: item.accrued,
									paid: item.paid,
									closingBalance: item.closingBalance,
									uploadDate: item.uploadDate,
								})
							}}
						>
							Изменить
						</button>

						<button
							type="button"
							onClick={() => deleteDebt(item.id)}
						>
							Удалить
						</button>
						<hr />
					</article>
				))}
			</section>

			<dialog className="debts--update" id="debts--update">
				<form id="form">
					<label>Введите номер лицевого счета</label>
					<input
						type="number"
						name="accountNumber"
						value={formData.accountNumber}
						onChange={handleChange}
					/><br />
					<label>Введите дату отчета</label>
					<input
						type="date"
						name="reportDate"
						value={formData.reportDate}
						onChange={handleChange}
					/><br />
					<label>Введите Начальное сальдо</label>
					<input
						type="number"
						name="openingBalance"
						value={formData.openingBalance}
						onChange={handleChange}
					/><br />
					<label>Введите Начисления сальдо</label>
					<input
						type="number"
						name="accrued"
						value={formData.accrued}
						onChange={handleChange}
					/><br />
					<label>Введите Платеж сальдо</label>
					<input
						type="number"
						name="paid"
						value={formData.paid}
						onChange={handleChange}
					/><br />
					<label>Введите Конечное сальдо сальдо</label>
					<input
						type="number"
						name="closingBalance"
						value={formData.closingBalance}
						onChange={handleChange}
					/><br />

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); updateDebt() }}
					>
						Отправить
					</button>

					<button
						type="submit"
						onClick={(e) => { e.preventDefault(); closeModule("debts--update") }}
					>
						Закрыть
					</button>
				</form>
			</dialog>
		</main>
	)
}


export default Debts