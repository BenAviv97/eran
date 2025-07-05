import { render, screen } from '@testing-library/react'
import Layout from '../components/Layout'

describe('Layout', () => {
  it('renders header', () => {
    render(
      <Layout>
        <p>child</p>
      </Layout>
    )
    expect(screen.getByText('LiveStreamPro')).toBeInTheDocument()
    expect(screen.getByText('child')).toBeInTheDocument()
  })
})
