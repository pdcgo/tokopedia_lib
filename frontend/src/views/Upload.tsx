/* eslint-disable react-hooks/exhaustive-deps */
import {
  CheckOutlined,
  DeleteOutlined,
  FilePptOutlined,
  SaveOutlined,
  UploadOutlined
} from '@ant-design/icons'
import {
  Button,
  Card,
  Checkbox,
  Divider,
  Input,
  Pagination,
  message
} from 'antd'
import React, { useEffect, useState } from 'react'
import { useRequest } from '../client'
import ProfileCard from '../components/ProfileCard'
import { Flex, FlexColumn } from '../styled_components'

export default function Upload (): React.ReactElement {
  const [query, setQuery] = useState({ page: 1, limit: 1, name: '' })
  const [showBottomPagination, setShowBottomPagination] = useState(false)
  const [messageApi, ctx] = message.useMessage()

  const { sender, response } = useRequest('GetTokopediaAkunList')

  useEffect(() => {
    sender({
      method: 'get',
      path: '/tokopedia/akun/list',
      params: {
        limit: query.limit,
        offset: (query.page - 1) * query.limit,
        search: query.name
      }
    })
  }, [query.limit, query.name, query.page])

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
          <Flex style={{ flex: 1 }}>
            <Input
              allowClear
              placeholder='Search Profile...'
              style={{ flex: 1 }}
              value={query.name}
              onChange={e =>
                setQuery(q => ({ ...q, page: 1, name: e.target.value }))
              }
            />
            <Button icon={<FilePptOutlined rev='paste' />}>Paste All</Button>
            <Button icon={<CheckOutlined rev='active' />}>Set Active</Button>
            <Button danger icon={<DeleteOutlined rev='remove' />}>
              Remove
            </Button>
            <Button
              style={{
                backgroundColor: '#f8da30',
                boxShadow: 'none',
                color: '#333'
              }}
              type='primary'
              icon={<SaveOutlined rev='save' />}
            >
              Save
            </Button>
            <Button
              type='primary'
              icon={<UploadOutlined rev='upload' />}
              style={{ boxShadow: 'none' }}
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
        {JSON.stringify(query)}
        <Pagination
          pageSize={query.limit}
          total={response?.pagination.count}
          showSizeChanger
          pageSizeOptions={[1, 2]}
          current={query.page}
          onChange={(page, size) => {
            if (query.limit !== size) {
              setQuery(q => ({ ...q, limit: size, page: 1 }))
            } else {
              setQuery(q => ({ ...q, limit: size, page }))
            }
          }}
        />
      </Flex>
      <div></div>
      {response?.data.map(profile => (
        <ProfileCard profile={profile} />
      ))}
      <div></div>
      {showBottomPagination && <Pagination pageSize={10} total={120} />}
    </FlexColumn>
  )
}
