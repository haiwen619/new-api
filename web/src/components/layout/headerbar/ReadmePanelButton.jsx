/*
Copyright (C) 2025 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/

import React, { useState } from 'react';
import { Button, Empty, SideSheet, Space, Spin, Typography } from '@douyinfe/semi-ui';
import { IconFile } from '@douyinfe/semi-icons';
import { API, showError } from '../../../helpers';
import { useIsMobile } from '../../../hooks/common/useIsMobile';
import MarkdownRenderer from '../../common/markdown/MarkdownRenderer';

const ReadmePanelButton = ({ t }) => {
  const isMobile = useIsMobile();
  const [visible, setVisible] = useState(false);
  const [loading, setLoading] = useState(false);
  const [loaded, setLoaded] = useState(false);
  const [content, setContent] = useState('');

  const loadReadme = async (force = false) => {
    if (loading) return;
    if (loaded && !force) return;

    setLoading(true);
    try {
      const res = await API.get('/api/readme/katu');
      const { success, data, message } = res.data || {};
      if (success) {
        setContent(data || '');
        setLoaded(true);
      } else {
        showError(message || t('加载说明文档失败'));
      }
    } catch (error) {
      showError(t('加载说明文档失败'));
    } finally {
      setLoading(false);
    }
  };

  const handleOpen = () => {
    setVisible(true);
    loadReadme();
  };

  return (
    <>
      <Button
        icon={<IconFile />}
        aria-label={t('说明文档')}
        onClick={handleOpen}
        theme='borderless'
        type='tertiary'
        className='!px-2.5 !text-current focus:!bg-semi-color-fill-1 dark:focus:!bg-gray-700 !rounded-full !bg-semi-color-fill-0 dark:!bg-semi-color-fill-1 hover:!bg-semi-color-fill-1 dark:hover:!bg-semi-color-fill-2'
      >
        <span className='hidden md:inline text-xs'>{t('说明文档')}</span>
      </Button>

      <SideSheet
        visible={visible}
        placement='right'
        width={isMobile ? '100%' : 760}
        title={
          <Space>
            <IconFile />
            <Typography.Text strong>{t('说明文档')}</Typography.Text>
          </Space>
        }
        bodyStyle={{ padding: 20, overflowY: 'auto' }}
        onCancel={() => setVisible(false)}
      >
        {loading && !content ? (
          <div className='w-full flex justify-center items-center py-16'>
            <Spin />
          </div>
        ) : content?.trim() ? (
          <MarkdownRenderer content={content} />
        ) : (
          <div className='w-full flex flex-col items-center gap-4 py-12'>
            <Empty description={t('暂无说明文档内容')} />
            <Button onClick={() => loadReadme(true)}>{t('重试')}</Button>
          </div>
        )}
      </SideSheet>
    </>
  );
};

export default ReadmePanelButton;

