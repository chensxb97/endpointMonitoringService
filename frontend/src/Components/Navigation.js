import React, { useState } from 'react'
import { Layout, Button } from 'antd'
import EndpointForm from './EndpointForm'
import StatusList from './StatusList'

const { Content } = Layout

export default function Navigation() {
    const Dashboard = 'Dashboard'
    const CreateEndpoint = 'Create Endpoint'

    const [active, setActive] = useState(Dashboard)

    return (
        <Layout >
            <Content style={{ padding: '24px' }}>
                <div style={{ marginBottom: '16px' }}>
                    <Button
                        type={active === Dashboard ? "primary" : ""}
                        onClick={() => setActive(Dashboard)}
                        style={{ marginRight: '8px' }}
                    >
                        Dashboard
                    </Button>
                    <Button
                        type={active === CreateEndpoint ? "primary" : ""}
                        onClick={() => setActive(CreateEndpoint)}
                    >
                        Create Endpoint
                    </Button>
                </div>
                {active === Dashboard && (
                    <StatusList />
                )}
                {active === CreateEndpoint && (
                    <EndpointForm />
                )}
            </Content>
        </Layout>
    )
}
