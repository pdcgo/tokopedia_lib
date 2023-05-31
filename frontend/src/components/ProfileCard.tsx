import {
  DeleteOutlined,
  ReloadOutlined,
  UploadOutlined,
  CopyOutlined,
  FilePptOutlined
} from '@ant-design/icons'
import { Card, Checkbox, Input, Select, Tooltip, Typography } from 'antd'
import React from 'react'
import { Flex, FlexColumn } from '../styled_components'

export type Profile = {
  readonly username: string
  readonly password: string
}

export default function ProfileCard (props: {
  profile: Profile,
  number: number
}): React.ReactElement {
  return (
    <Card
      title={
        <Checkbox style={{ userSelect: 'none' }}>
          {props.number + '. '}
          {props.profile.username}
        </Checkbox>
      }
      hoverable
      size='small'
      type='inner'
      actions={[
        <Tooltip title='Upload' placement='bottom' showArrow={false}>
          <UploadOutlined
            style={{ color: '#058b3d' }}
            rev={'upload'}
            key='upload'
          />
        </Tooltip>,
        <Tooltip title='Copy' placement='bottom' showArrow={false}>
          <CopyOutlined style={{ color: '#db17cb' }} rev={'copy'} key='copy' />
        </Tooltip>,
        <Tooltip title='Paste' placement='bottom' showArrow={false}>
          <FilePptOutlined
            style={{ color: '#172bdb' }}
            rev={'paste'}
            key='paste'
          />
        </Tooltip>,
        <Tooltip title='Reset' placement='bottom' showArrow={false}>
          <ReloadOutlined
            style={{ color: '#707070' }}
            rev={'reset'}
            key='reset'
          />
        </Tooltip>,
        <Tooltip title='Remove' placement='bottom' showArrow={false}>
          <DeleteOutlined
            style={{ color: '#f2113a' }}
            rev={'delete'}
            key='delete'
          />
        </Tooltip>
      ]}
    >
      <Flex style={{ width: '100%' }}>
        <FlexColumn style={{ flex: 1 }}>
          <FlexColumn style={{ rowGap: '5px' }}>
            <Typography.Text>Username :</Typography.Text>
            <Input value={props.profile.username} placeholder='username' />
          </FlexColumn>
          <FlexColumn style={{ rowGap: '5px' }}>
            <Typography.Text>Password :</Typography.Text>
            <Input.Password
              value={props.profile.password}
              placeholder='⁎⁎⁎⁎⁎⁎⁎⁎'
            />
          </FlexColumn>
          <FlexColumn style={{ rowGap: '5px' }}>
            <Typography.Text>Upload Limit :</Typography.Text>
            <Input placeholder='1000' />
          </FlexColumn>
          <div></div>
          <Checkbox style={{ userSelect: 'none' }}>Active</Checkbox>
        </FlexColumn>
        <FlexColumn style={{ flex: 1 }}>
          <FlexColumn style={{ rowGap: '5px' }}>
            <Typography.Text>Markup :</Typography.Text>
            <Select placeholder='Choose Markup Data' />
          </FlexColumn>
          <FlexColumn style={{ rowGap: '5px' }}>
            <Typography.Text>Spin :</Typography.Text>
            <Select placeholder='Choose Spin Data' />
          </FlexColumn>
          <FlexColumn style={{ rowGap: '5px' }}>
            <Typography.Text>Collection :</Typography.Text>
            <Select placeholder='Choose Collection Data' />
          </FlexColumn>
        </FlexColumn>
      </Flex>
    </Card>
  )
}
