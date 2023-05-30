/* eslint-disable @typescript-eslint/no-non-null-assertion */
import {
  ArrowLeftOutlined,
  CloudUploadOutlined,
  ShopOutlined
} from '@ant-design/icons'
import { Button, Card, Tabs } from 'antd'
import { useEffect, useState } from 'react'
import {
  AppContainer,
  AppContainer2,
  BackButtonContainer,
  FlexColumn,
  FloatingMenu
} from './styled_components'
import AddAccount from './views/AddAccount'
import Upload from './views/Upload'

const menus = [
  {
    key: 'add_account',
    name: 'Add Account',
    icon: ShopOutlined,
    child: AddAccount
  },
  {
    key: 'upload',
    name: 'Upload',
    icon: CloudUploadOutlined,
    child: Upload
  }
]

function App () {
  const [activeMenu, setActiveMenu] = useState('add_account')
  const [showFloatMenu, setShowFloatMenu] = useState(false)

  useEffect(() => {
    const observer = new IntersectionObserver(
      entry => {
        console.log(entry[0].isIntersecting)
        if (!entry[0].isIntersecting) setShowFloatMenu(true)
        else setShowFloatMenu(false)
      },
      { threshold: [0] }
    )

    observer.observe(document.getElementById('add_account')!)

    return () => {
      observer.unobserve(document.getElementById('add_account')!)
    }
  }, [])

  return (
    <AppContainer>
      <BackButtonContainer>
        <Button
          type='default'
          size='large'
          shape='circle'
          style={{ fontSize: 18 }}
          icon={<ArrowLeftOutlined rev="0" />}
        ></Button>
      </BackButtonContainer>
      <AppContainer2>
        <FloatingMenu show={showFloatMenu}>
          <Card size='small' title='Recent Page' type='inner'>
            <FlexColumn>
              {menus.map(menu => (
                <Button
                  key={menu.key}
                  type={activeMenu == menu.key ? 'primary' : 'default'}
                  onClick={() => setActiveMenu(menu.key)}
                  icon={<menu.icon rev={menu.key} />}
                >
                  {menu.name}
                </Button>
              ))}
            </FlexColumn>
          </Card>
        </FloatingMenu>
        <Tabs
          style={{ height: '100%' }}
          activeKey={activeMenu}
          onChange={setActiveMenu}
          items={menus.map(menu => ({
            key: menu.key,
            label: (
              <span
                id={menu.key}
                style={{ fontWeight: 400, userSelect: 'none' }}
              >
                <menu.icon rev={menu.key} />
                {menu.name}
              </span>
            ),
            children: <menu.child />
          }))}
        />
      </AppContainer2>
    </AppContainer>
  )
}

export default App
