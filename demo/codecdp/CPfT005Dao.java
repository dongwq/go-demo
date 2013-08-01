package com.cdp.xauth.dao;

import java.util.List;
import java.util.Map;

import org.mybatis.spring.support.SqlSessionDaoSupport;
import org.springframework.stereotype.Repository;

import com.cdp.xauth.model.bo.{{.ClassName}};
import com.cdp.xauth.model.bo.{{.ClassName}}Example;
import com.cdp.xauth.model.mapper.{{.ClassName}}Mapper;

@Repository
public class {{.ClassName}}Dao extends SqlSessionDaoSupport implements {{.ClassName}}Mapper {
	
	String namespace = "com.cdp.xauth.model.mapper.{{.ClassName}}Mapper.";
	@Override
	public int countByExample({{.ClassName}}Example example) {
		return this.getSqlSession().getMapper({{.ClassName}}Mapper.class).countByExample(example);
	}

	@Override
	public int deleteByExample({{.ClassName}}Example example) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public int deleteByPrimaryKey(Integer userId) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public int insert({{.ClassName}} record) {
		return this.getSqlSession().getMapper({{.ClassName}}Mapper.class).insert(record);
	}

	@Override
	public int insertSelective({{.ClassName}} record) {
		return this.getSqlSession().getMapper({{.ClassName}}Mapper.class).insertSelective(record);
	}

	@Override
	public List<{{.ClassName}}> selectByExample({{.ClassName}}Example example) {
		return this.getSqlSession().getMapper({{.ClassName}}Mapper.class).selectByExample(example);
	}

	@Override
	public {{.ClassName}} selectByPrimaryKey(Integer userId) {
		return this.getSqlSession().getMapper({{.ClassName}}Mapper.class).selectByPrimaryKey(userId);
	}

	@Override
	public int updateByExampleSelective({{.ClassName}} record, {{.ClassName}}Example example) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public int updateByExample({{.ClassName}} record, {{.ClassName}}Example example) {
		// TODO Auto-generated method stub
		return 0;
	}

	@Override
	public int updateByPrimaryKeySelective({{.ClassName}} record) {
		return this.getSqlSession().getMapper({{.ClassName}}Mapper.class).updateByPrimaryKeySelective(record);
	}

	@Override
	public int updateByPrimaryKey({{.ClassName}} record) {
		return this.getSqlSession().getMapper({{.ClassName}}Mapper.class).updateByPrimaryKey(record);
	}
	
	public List<Map> selectByExampleForMap({{.ClassName}}Example example) {
		return this.getSqlSession().selectList(namespace+"selectByExampleForMap", example);
	}
	
	public List<Map> getAuthSysByUserId(Integer userid){
		return this.getSqlSession().selectList(namespace+"getAuthSysByUserId",userid);
	}
}
