import {
  CheckOutlined,
  DeleteOutlined,
  FilePptOutlined,
  UploadOutlined
} from '@ant-design/icons'
import { Button, Card, Checkbox, Divider, Pagination, message } from 'antd'
import React, { useEffect, useState } from 'react'
import type { Profile } from '../components/ProfileCard'
import ProfileCard from '../components/ProfileCard'
import { Flex, FlexColumn } from '../styled_components'

const profiles: Profile[] = [
  {
    password: '',
    username: 'nama_toko'
  },
  {
    password: '',
    username: 'nama_toko'
  },
  {
    password: '',
    username: 'nama_toko'
  },
  {
    password: '',
    username: 'nama_toko'
  }
  //   {
  //     password: '',
  //     username: 'nama_toko'
  //   },
  //   {
  //     password: '',
  //     username: 'nama_toko'
  //   },
  //   {
  //     password: '',
  //     username: 'nama_toko'
  //   },
  //   {
  //     password: '',
  //     username: 'nama_toko'
  //   },
  //   {
  //     password: '',
  //     username: 'nama_toko'
  //   },
  //   {
  //     password: '',
  //     username: 'nama_toko'
  //   }
]

export default function Upload (): React.ReactElement {
  const [showBottomPagination, setShowBottomPagination] = useState(false)
  const [messageApi, ctx] = message.useMessage()

  useEffect(() => {
    const observer = new IntersectionObserver(
      function (entry) {
        if (!entry[0].isIntersecting) setShowBottomPagination(true)
        else setShowBottomPagination(false)
      },
      { threshold: [0] }
    )

    // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
    observer.observe(document.getElementById('top-pagination')!)

    return () => {
      // eslint-disable-next-line @typescript-eslint/no-non-null-assertion
      observer.unobserve(document.getElementById('top-pagination')!)
    }
  }, [])

  return (
    <FlexColumn>
      {ctx}
      <Card size='small' title='Setting Tokopedia Upload'>
        <Flex
          style={{
            justifyContent: 'space-between',
            alignItems: 'center'
          }}
        >
          <Checkbox>Select All</Checkbox>
          <Flex>
            <Button icon={<FilePptOutlined rev='paste' />}>Paste All</Button>
            <Button icon={<CheckOutlined rev='active' />}>Set Active</Button>
            <Button danger icon={<DeleteOutlined rev='remove' />}>
              Remove
            </Button>
            <Button
              type='primary'
              icon={<UploadOutlined rev='upload' />}
              onClick={() => {
                messageApi.open({
                  type: 'loading',
                  content: 'Running process, please wait...',
                  duration: 0,
                  key: 'loading'
                })

                setTimeout(() => {
                  messageApi.destroy('loading')
                }, 5000)
              }}
            >
              Start Upload
            </Button>
          </Flex>
        </Flex>
      </Card>
      <Divider dashed style={{ margin: '5px 0' }} />
      <Flex style={{ justifyContent: 'flex-start' }} id='top-pagination'>
        <Pagination pageSize={10} total={120} />
      </Flex>
      <div></div>
      {profiles.map(profile => (
        <ProfileCard profile={profile} />
      ))}
      <div></div>
      {showBottomPagination && <Pagination pageSize={10} total={120} />}
    </FlexColumn>
  )
}
