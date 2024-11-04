import React, { useState, useEffect, useMemo } from 'react';
import { Menu, Modal, Input, Button } from 'antd';
import { useLocation, useNavigate } from 'react-router-dom';
import './Sidebar.scss';

type MenuItem = {
  key: string;
  icon?: React.ReactNode;
  children?: MenuItem[];
  label: React.ReactNode;
};

const getItem = (
  label: React.ReactNode,
  key: string,
  icon?: React.ReactNode,
  children?: MenuItem[]
): MenuItem => ({
  key,
  icon,
  children,
  label,
});

const menuList: MenuItem[] = [
  { key: '/', label: '首页', icon: '' },
  {
    key: '/dashboard',
    label: '用户仪表盘',
    icon: '',
    children: [
      { key: '/parking-records', label: '停车记录', icon: '' },
      { key: '/create-vehicle', label: '绑定车辆', icon: '' },
      { key: '/vehicle-info', label: '车辆信息', icon: '' },
      { key: '/violations', label: '违章停车', icon: '' },
      { key: '/bookings', label: '停车位预约', icon: '' },
      { key: '/payments', label: '支付', icon: '' },
    ],
  },
  { key: '/parking-space', label: '空闲车位', icon: '' },
  { key: '/parking-lots', label: '停车场列表', icon: '' },
  { key: '/user', label: '用户页面', icon: '' },
  { key: '/admin-dashboard', label: '管理员面板', icon: '' },
];

const Sidebar: React.FC = () => {
  const path = useLocation().pathname;
  const navigate = useNavigate();
  const [openKeys, setOpenKeys] = useState<string[]>([]);
  const [passwordModalVisible, setPasswordModalVisible] = useState(false);
  const [password, setPassword] = useState('');

  const tempMenuList = useMemo(() => {
    const buildMenu = (list: MenuItem[]): MenuItem[] =>
      list.map(({ key, label, icon, children }) => ({
        key,
        icon,
        label,
        children: children ? buildMenu(children) : undefined,
      }));
    return buildMenu(
      menuList.map((item) => getItem(item.label, item.key, item.icon, item.children))
    );
  }, []);

  useEffect(() => {
    // 这里不再需要更新 tempPath，因为我们直接使用 navigate 进行跳转
  }, [path]);

  const onOpenChange: Menu['onOpenChange'] = (keys) => {
    setOpenKeys(keys);
  };

  const handleMenuClick: Menu['onClick'] = ({ key }) => {
    if (key === '/admin-dashboard') {
      setPasswordModalVisible(true); // 显示密码输入模态框
    } else {
      navigate(key); // 对于其他菜单项，直接跳转
    }
  };

  const handlePasswordSubmit = () => {
    if (password === 'root') {
      navigate('/admin-dashboard'); // 密码正确，跳转到管理员面板
      setPasswordModalVisible(false); // 关闭密码输入模态框
    } else {
      // 密码错误，可以在这里添加错误提示
      console.log('密码错误');
    }
  };

  return (
    <div className="sidebar">
      <h2>菜单</h2>
      <Menu
        onClick={handleMenuClick}
        onOpenChange={onOpenChange}
        mode="vertical"
        inlineIndent={24}
        selectedKeys={[path]} // 使用当前路径作为选中的菜单项
        openKeys={openKeys}
        items={tempMenuList}
      />
      {passwordModalVisible && (
        <Modal
          title="请输入密码"
          visible={passwordModalVisible}
          onCancel={() => setPasswordModalVisible(false)}
          footer={null}
        >
          <Input
            type="password"
            className="password-input"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="请输入密码"
          />
          <br />
          <br />
          <Button className="submit-button" type="primary" onClick={handlePasswordSubmit}>
            提交
          </Button>
        </Modal>
      )}
    </div>
  );
};

export default Sidebar;
