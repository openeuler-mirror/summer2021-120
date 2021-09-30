#!/usr/local/easyops/python/bin/python
# -*- coding: utf-8 -*-

"""
功能：通用化jenkins构建，可传1个任务名、0或多个参数值
设计：针对以前的版本，修改为使用jenkinsapi.jenkins中的Jenkins
参数：
    JENKINS_SERVER    构建服务地址
    JENKINS_USER      构建账号名称
    JENKINS_PASSWORD  构建账号密码
    JOB_NAMES         构建任务名称
    PARAMS            构建任务参数，构建作业需要的参数，如name:admin;version:1123
    TIMEOUT           构建超时时间
"""
import logging
import sys
import requests

from jenkinsapi.jenkins import Jenkins
import os
import time
import requests.utils

sys.path.append("libs")
reload(sys)
sys.setdefaultencoding('utf-8')

logging.basicConfig(format='[%(asctime)s] %(message)s')
log = logging.getLogger(name='log')
log.setLevel(logging.DEBUG)

CI_BUILD_ID = os.environ.get('CI_BUILD_ID')
log.debug('【CI_BUILD_ID】: %s' % CI_BUILD_ID)

CI_WORKSPACE = os.environ.get('CI_WORKSPACE')
log.debug('【CI_WORKSPACE】: %s' % CI_WORKSPACE)

URL = os.environ.get('PLUGIN_URL')
log.debug('【Jenkins】: %s' % URL)

USER = os.environ.get('PLUGIN_USER')
log.debug('【USER】: %s' % USER)

PASSWORD = os.environ.get('PLUGIN_PASSWORD')
log.debug('【PASSWORD】: %s' % "********")

TIMEOUT = os.environ.get('PLUGIN_TIMEOUT', 600)
log.debug('【TIMEOUT】: %s' % TIMEOUT)

JOB_NAME = os.environ.get('PLUGIN_JOB_NAME')
log.debug('【JOB_NAME】: %s' % JOB_NAME)

PARAMS = os.environ.get('PLUGIN_PARAMS')
log.debug('【PARAMS】: %s' % PARAMS)

def get_jenkins():
    try:
        log.debug('start to connect to jenkins server')
        jenkins_server = Jenkins(baseurl=URL, username=USER, password=PASSWORD)
        return jenkins_server
    except Exception as err:
        log.error("Can not connect to the jenkins_server: %s, %s" % (URL, err))
        exit(code=1)


def build_job():
    jks = get_jenkins()
    job = jks.get_job(jobname=JOB_NAME)

    log.debug('start to build job【%s】, params is: %s' % (job.name,  PARAMS))
    jks.build_job(jobname=JOB_NAME, params=PARAMS)

    time.sleep(10)

    build_no = job.get_last_buildnumber()
    if build_no in job.get_build_dict():
        build = job.get_last_build()
        build = job.get_build(buildnumber=build_no)  # 为什么再调用一次？日志打印后，status不会更新，build.get_status()为空
        if not build.is_good():
            exit(code=1)
    else:
        log.debug('last build is not existed')
        exit(code=1)


if __name__ == '__main__':
    build_job()
