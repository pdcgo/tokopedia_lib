import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import { ConfigProvider, ThemeConfig } from 'antd'
import './index.scss'
import 'antd/dist/reset.css'

const theme: ThemeConfig = {
  token: {
    fontFamily:
      "Roboto, -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Helvetica Neue', Arial, 'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji' ",
    fontSize: 13.4,
    colorPrimary: '#106C36',
    colorInfo: '#4fc078',
    colorError: '#e7b072',
    colorErrorHover: '#e0732b',
    colorErrorBorder: '#e7b072',
    colorErrorBorderHover: '#d88b34'
  },
  components: {
    Tooltip: {
      colorBgDefault: '#4e4e4e'
    },
    Pagination: {
      colorBgContainer: '#fefffece'
    },
    Input: {
      fontSize: 14,
      fontSizeLG: 13.5,
    },
    Button: {
      boxShadow: 'none',
      boxShadowSecondary: 'none',
      boxShadowTertiary: 'none'
    }
  }
}

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <ConfigProvider theme={theme}>
    <App />
  </ConfigProvider>
)
