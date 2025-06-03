import React from "react"
import axios from "axios"
import { Card, Form, Input, Button, Space } from "antd"
import { MinusCircleOutlined, PlusOutlined } from "@ant-design/icons"

const backendUrl = process.env.REACT_APP_BACKEND_URL

const EndpointForm = () => {
    const [form] = Form.useForm()

    const onFinish = async (values) => {
        try {
            const response = await axios.post(`${backendUrl}/targets/create`, {
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    endpoint: values.endpoint,
                    labels: values.labels || [],
                })
            })
            const data = await response.json()
            console.log("Endpoint created successfully:", data)
            form.resetFields()
        } catch (error) {
            console.error("Error creating endpoint:", error)
        }
    }

    return (
        <Card style={{ maxWidth: 600, maxHeight: 500, margin: "0 auto" }}>
            <Form form={form} onFinish={onFinish} layout="vertical" autoComplete="off">
                <Form.Item
                    label="Endpoint"
                    name="endpoint"
                    rules={[{ required: true, message: "Please input the endpoint!" }]}
                >
                    <Input placeholder="https://example.com" />
                </Form.Item>

                <Form.List name="labels">
                    {(fields, { add, remove }) => (
                        <>
                            <label>Labels</label>
                            {fields.map(({ key, name, ...restField }) => (
                                <Space key={key} style={{ display: "flex", marginBottom: 8 }} align="baseline">
                                    <Form.Item
                                        {...restField}
                                        name={[name, "key"]}
                                        rules={[{ required: true, message: "Missing key" }]}
                                    >
                                        <Input placeholder="Key" />
                                    </Form.Item>
                                    <Form.Item
                                        {...restField}
                                        name={[name, "value"]}
                                        rules={[{ required: true, message: "Missing value" }]}
                                    >
                                        <Input placeholder="Value" />
                                    </Form.Item>
                                    <MinusCircleOutlined onClick={() => remove(name)} />
                                </Space>
                            ))}
                            <Form.Item>
                                <Button type="dashed" onClick={() => add()} block icon={<PlusOutlined />}>
                                    Add Label
                                </Button>
                            </Form.Item>
                        </>
                    )}
                </Form.List>

                {/* Submit */}
                <Form.Item>
                    <Button type="primary" htmlType="submit">
                        Submit Endpoint
                    </Button>
                </Form.Item>
            </Form>
        </Card>
    )
}

export default EndpointForm
