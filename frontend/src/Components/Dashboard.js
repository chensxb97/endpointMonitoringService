import React, { useEffect, useState } from "react"
import Timer from "./Timer"
import axios from "axios"
import { Card } from "antd"

const backendUrl = process.env.REACT_APP_BACKEND_URL
const timeLimit = 5000

const Dashboard = () => {
    const [statuses, setStatuses] = useState([])
    const [loading, setLoading] = useState(true)
    const [animateKey, setAnimateKey] = useState(0)

    const fetchStatuses = async () => {
        try {
            const response = await axios.get(`${backendUrl}/statuses`)
            if (response?.data) {
                setStatuses(response.data)
            }
        } catch (error) {
            console.error("Error fetching statuses:", error)
        }
        setLoading(false)
    }

    useEffect(() => {
        fetchStatuses()
        const intervalId = setInterval(() => {
            fetchStatuses()
            setAnimateKey((prevKey) => !prevKey)
        }, timeLimit)
        return () => clearInterval(intervalId)
        // eslint-disable-next-line
    }, [])

    return (
        <Card style={{ maxWidth: 600, maxHeight: 600, margin: "0 auto" }}>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <>
                    <Timer customTime={timeLimit} />
                    <div style={{ display: "flex", justifyContent: "center", marginTop: "20px" }}>
                        <table style={{ borderCollapse: "collapse", border: "3px solid grey" }}>
                            <thead>
                                <tr>
                                    <th style={{ padding: "10px", border: "2px solid #ddd" }}>Endpoint</th>
                                    <th style={{ padding: "10px", border: "2px solid #ddd" }}>Module</th>
                                    <th style={{ padding: "10px", border: "2px solid #ddd" }}>Custom Labels</th>
                                    <th style={{ padding: "10px", border: "2px solid #ddd" }}>Status</th>
                                </tr>
                            </thead>
                            <tbody>
                                {statuses.map((status) => (
                                    <tr key={status?.endpoint}>
                                        <td style={{ padding: "15px", border: "2px solid #ddd" }}>
                                            {status?.endpoint}
                                        </td>
                                        <td style={{ padding: "15px", border: "2px solid #ddd" }}>
                                            {status?.module}
                                        </td>
                                        <td style={{ padding: "15px", border: "2px solid #ddd" }}>
                                            {status?.labels?.length > 0 ? (
                                                status.labels.map((label, index) => (
                                                    <div key={index}>
                                                        {label.key}{` -> `}{label.value}
                                                    </div>
                                                ))
                                            ) : (
                                                <div>No custom labels</div>
                                            )}
                                        </td>
                                        <td
                                            key={animateKey}
                                            style={{
                                                padding: "15px",
                                                border: "2px solid #ddd",
                                                color: status?.status === "live" ? "green" : "red",
                                                animation: status?.status ? "bounceOut 1s ease" : "",
                                            }}
                                        >
                                            {status?.status?.toUpperCase()}
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </>
            )
            }
        </Card>)
}

export default Dashboard