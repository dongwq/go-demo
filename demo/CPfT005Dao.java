package com.cdp.xauth.dao;
{{.UserName}}
import java.util.List;
import java.util.Map;

import org.mybatis.spring.support.SqlSessionDaoSupport;
import org.springframework.stereotype.Repository;

import com.cdp.xauth.model.bo.CPfT005;
import com.cdp.xauth.model.bo.CPfT005Example;
import com.cdp.xauth.model.mapper.CPfT005Mapper;

@Repository
public class CPfT005Dao extends SqlSessionDaoSupport implements CPfT005Mapper {
	
	String namespace = "com.cdp.xauth.model.mapper.CPfT005Mapper.";
	@Override
	public int countByExample(CPfT005Example example) {
		return this.getSqlSession().getMapper(CPfT005Mapper.class).countByExample(example);
	}

	@Override
	public int deleteByExample(CPfT005Example example) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public int deleteByPrimaryKey(Integer userId) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public int insert(CPfT005 record) {
		return this.getSqlSession().getMapper(CPfT005Mapper.class).insert(record);
	}

	@Override
	public int insertSelective(CPfT005 record) {
		return this.getSqlSession().getMapper(CPfT005Mapper.class).insertSelective(record);
	}

	@Override
	public List<CPfT005> selectByExample(CPfT005Example example) {
		return this.getSqlSession().getMapper(CPfT005Mapper.class).selectByExample(example);
	}

	@Override
	public CPfT005 selectByPrimaryKey(Integer userId) {
		return this.getSqlSession().getMapper(CPfT005Mapper.class).selectByPrimaryKey(userId);
	}

	@Override
	public int updateByExampleSelective(CPfT005 record, CPfT005Example example) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public int updateByExample(CPfT005 record, CPfT005Example example) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public int updateByPrimaryKeySelective(CPfT005 record) {
		return this.getSqlSession().getMapper(CPfT005Mapper.class).updateByPrimaryKeySelective(record);
	}

	@Override
	public int updateByPrimaryKey(CPfT005 record) {
		return this.getSqlSession().getMapper(CPfT005Mapper.class).updateByPrimaryKey(record);
	}
	
	public List<Map> selectByExampleForMap(CPfT005Example example) {
		return this.getSqlSession().selectList(namespace+"selectByExampleForMap", example);
	}
	
	public List<Map> getAuthSysByUserId(Integer userid){
		return this.getSqlSession().selectList(namespace+"getAuthSysByUserId",userid);
	}
}
