import React, { useState } from 'react'
import { Layout, Button } from 'antd'
import EndpointForm from './EndpointForm'
import Dashboard from './Dashboard'

const { Content } = Layout

const DASHBOARD = 'Dashboard'
const CREATE_ENDPOINT = 'Create Endpoint'

export default function Navigation() {
    const [active, setActive] = useState(DASHBOARD)

    return (
        <Layout >
            <Content style={{ padding: '24px' }}>
                <div style={{ marginBottom: '16px' }}>
                    <Button
                        type={active === DASHBOARD ? "primary" : ""}
                        onClick={() => setActive(DASHBOARD)}
                        style={{ marginRight: '8px' }}
                    >
                        {DASHBOARD}
                    </Button>
                    <Button
                        type={active === CREATE_ENDPOINT ? "primary" : ""}
                        onClick={() => setActive(CREATE_ENDPOINT)}
                    >
                        {CREATE_ENDPOINT}
                    </Button>
                </div>
                {active === DASHBOARD && (
                    <Dashboard />
                )}
                {active === CREATE_ENDPOINT && (
                    <EndpointForm />
                )}
            </Content>
        </Layout>
    )
}
