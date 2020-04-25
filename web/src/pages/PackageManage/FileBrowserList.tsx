import {AnyAction, Dispatch} from 'redux';
import {Card, Form, Table} from 'antd';
import React, {PureComponent} from 'react';

import {FormComponentProps} from 'antd/lib/form';
import {connect} from 'dva';
import {ConnectState, FileBrowserModelState} from '@/models/connect';
import PageHeaderLayout from '../../layouts/PageHeaderLayout';
import {formatDate, formatSize} from '../../utils/utils';
import {FileOutlined, FolderOutlined, RollbackOutlined} from "@ant-design/icons/lib";

export interface FileBrowserListProps extends FormComponentProps {
  dispatch: Dispatch<AnyAction>;
  user: FileBrowserModelState;
  loading: any;
}

export interface FileBrowserListState {
  folder: string,
}

@connect((state: ConnectState) => ({
  loading: state.loading.models.user,
  user: state.fileBrowser,
}))
class FileBrowserList extends PureComponent<FileBrowserListProps, FileBrowserListState> {
  constructor(props: FileBrowserListProps) {
    super(props);
    this.state = {
      folder: '/',
    };
  }

  componentDidMount() {
    this.listFolder("/")
  }

  listFolder(basePath: string) {
    this.dispatch({
      type: 'fileBrowser/fetch',
      search: {'basePath': basePath},
    });
    this.setState({
      folder: basePath
    })
  }

  upperFolder(currentPath: string) {
    let index = currentPath.lastIndexOf("/");
    if (index > 0) {
      let targetPath = currentPath.substring(0, index);
      this.listFolder(targetPath);
    } else {
      this.listFolder("/");
    }

  }

  dispatch = (action: any) => {
    const {dispatch} = this.props;
    dispatch(action);
  };


  render() {
    const {
      loading,
      user: {data},
    } = this.props;

    const {list} = data || {};

    const columns = [
      {
        title: '名称',
        dataIndex: 'name',
        render: (val: any, record: any) => {
          return (<div>
            {record.dir &&
            <FolderOutlined style={{color: '#f7a517'}}/>
            }
            {
              !record.dir &&
              <FileOutlined style={{color: '#000000'}}/>
            }
            <a>&nbsp;{val}</a>
          </div>)
        },
      },
      {
        title: '大小',
        dataIndex: 'size',
        render: (val: any) => <span>{formatSize(val)}</span>,
      },
      {
        title: '更新时间',
        dataIndex: 'modifyTime',
        render: (val: any) => <span>{formatDate(val, 'YYYY-MM-DD HH:mm')}</span>,
      },
    ];


    return (
      <PageHeaderLayout title="文件管理">
        <Card bordered={false}>
          <div>
            <div
              style={{marginBottom: '8px', marginLeft: '18px'}}>
              <RollbackOutlined onClick={this.upperFolder.bind(this, this.state.folder)} style={{color: '#f7a517'}}/>
              <span style={{marginLeft: '8px'}}> {this.state.folder}</span>
            </div>
            <Table
              loading={loading}
              rowKey={(record: any) => record.id}
              dataSource={list}
              columns={columns}
              pagination={false}
              size="small"
              onRow={record => {
                return {
                  onClick: event => {
                    if (record.dir) {
                      this.listFolder(record.id)
                    }
                  }, // 点击行
                  onMouseEnter: event => {
                  }, // 鼠标移入行
                  onMouseLeave: event => {
                  },
                };
              }}
            />
          </div>
        </Card>
      </PageHeaderLayout>
    );
  }
}

export default Form.create<FileBrowserListProps>()(FileBrowserList);
