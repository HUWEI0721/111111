<template>
  <div class="history-container" :class="roleClass">
    <!-- 用户信息卡片 -->
    <el-card class="user-info-card" shadow="hover">
      <div class="user-info">
        <div class="info-header">
          <h2>历史记录查询</h2>
          <div class="role-icon">
            <i :class="roleIcon"></i>
          </div>
        </div>
        <div class="info-content">
          <div class="info-item">
            <i class="el-icon-user"></i>
            <span class="user-text">当前用户: <strong>{{ name }}</strong></span>
          </div>
          <el-divider direction="vertical"></el-divider>
          <div class="info-item">
            <i class="el-icon-s-custom"></i>
            <span class="role-text">用户角色: <strong>{{ userType }}</strong></span>
          </div>
        </div>
      </div>
    </el-card>

    <!-- 数据列表卡片 -->
    <el-card class="main-data-card" shadow="hover">
      <el-table v-loading="loading" :data="tracedata" style="width: 100%" class="custom-table">
        <el-table-column type="expand">
          <template slot-scope="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <!-- 茅台酒信息 -->
              <div class="info-section">
                <div class="section-header">
                  <i class="el-icon-ship"></i>
                  <span class="trace-text" style="color: #67C23A;">茅台酒信息</span>
                </div>
                <el-row :gutter="20">
                  <el-col :span="12">
                    <el-form-item label="茅台酒名称：">
                      <span>{{ props.row.brewer_input.moutai_name }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="产地：">
                      <span>{{ props.row.brewer_input.moutai_origin }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="酿造开始时间：">
                      <span>{{ props.row.brewer_input.moutai_start }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="酿造完成时间：">
                      <span>{{ props.row.brewer_input.moutai_end }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="酿酒师名称：">
                      <span>{{ props.row.brewer_input.brewer_name }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </div>

              <!-- 酿造厂信息 -->
              <div class="info-section">
                <div class="section-header">
                  <i class="el-icon-office-building"></i>
                  <span class="trace-text" style="color: #409EFF;">酿造厂信息</span>
                </div>
                <el-row :gutter="20">
                  <el-col :span="12">
                    <el-form-item label="商品名称：">
                      <span>{{ props.row.distillery_input.moutai_productName }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="酿造批次：">
                      <span>{{ props.row.distillery_input.moutai_batch }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="酿造时间：">
                      <span>{{ props.row.distillery_input.moutai_productionTime }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="酿造厂名称：">
                      <span>{{ props.row.distillery_input.distillery_name }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="联系电话：">
                      <span>{{ props.row.distillery_input.distillery_phone }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </div>

              <!-- 物流信息 -->
              <div class="info-section">
                <div class="section-header">
                  <i class="el-icon-truck"></i>
                  <span class="trace-text" style="color: #E6A23C;">物流轨迹信息</span>
                </div>
                <el-row :gutter="20">
                  <el-col :span="12">
                    <el-form-item label="司机姓名：">
                      <span>{{ props.row.driver_input.driver_name }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="车厢温度：">
                      <span>{{ props.row.driver_input.driver_temperature }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="联系电话：">
                      <span>{{ props.row.driver_input.driver_phone }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="车牌号：">
                      <span>{{ props.row.driver_input.driver_carNumber }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="运输工具：">
                      <span>{{ props.row.driver_input.driver_transport }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </div>

              <!-- 销售终端信息 -->
              <div class="info-section">
                <div class="section-header">
                  <i class="el-icon-shopping-cart-full"></i>
                  <span class="trace-text" style="color: #909399;">销售终端信息</span>
                </div>
                <el-row :gutter="20">
                  <el-col :span="12">
                    <el-form-item label="存入时间：">
                      <span>{{ props.row.shop_input.shop_storeTime }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="销售时间：">
                      <span>{{ props.row.shop_input.shop_sellTime }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="销售点名称：">
                      <span>{{ props.row.shop_input.shop_shopName }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="销售点位置：">
                      <span>{{ props.row.shop_input.shop_shopAddress }}</span>
                    </el-form-item>
                  </el-col>
                  <el-col :span="12">
                    <el-form-item label="销售点电话：">
                      <span>{{ props.row.shop_input.shop_shopPhone }}</span>
                    </el-form-item>
                  </el-col>
                </el-row>
              </div>
            </el-form>
          </template>
        </el-table-column>

        <el-table-column label="溯源码" prop="traceability_code" align="center" min-width="180">
          <template slot-scope="scope">
            <span class="trace-code">{{ scope.row.traceability_code }}</span>
          </template>
        </el-table-column>
        <el-table-column label="茅台酒名称" prop="brewer_input.moutai_name" align="center" min-width="150">
          <template slot-scope="scope">
            <el-tag :type="getTagType(scope.row.brewer_input.moutai_name)" effect="dark">
              {{ scope.row.brewer_input.moutai_name }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column label="酿造完成时间" prop="brewer_input.moutai_end" align="center" min-width="180">
          <template slot-scope="scope">
            <i class="el-icon-time"></i>
            <span style="margin-left: 5px">{{ scope.row.brewer_input.moutai_end }}</span>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getMoutaiInfo, getMoutaiList, getAllMoutaiInfo, getMoutaiHistory } from '@/api/trace'

export default {
  name: 'History',
  data() {
    return {
      tracedata: [],
      loading: false,
      input: ''
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ]),
    roleClass() {
      switch (this.userType) {
        case '茅台厂商': return 'fisherman-bg'
        case '酿造厂': return 'factory-bg'
        case '物流司机信息': return 'driver-bg'
        case '销售终端': return 'shop-bg'
        default: return 'default-bg'
      }
    },
    roleIcon() {
      switch (this.userType) {
        case '茅台厂商': return 'el-icon-ship'
        case '酿造厂': return 'el-icon-office-building'
        case '物流司机信息': return 'el-icon-truck'
        case '销售终端': return 'el-icon-shopping-cart-full'
        default: return 'el-icon-user'
      }
    }
  },
  created() {
    getMoutaiList().then(res => {
      this.tracedata = JSON.parse(res.data)
    })
  },
  methods: {
    getTagType(name) {
      const types = ['', 'success', 'warning', 'danger', 'info']
      return types[Math.abs(name.length) % types.length]
    }
  }
}
</script>

<style lang="scss" scoped>
.history-container {
  padding: 30px;
  min-height: 100vh;
  transition: all 0.3s ease;

  &.fisherman-bg {
    background: linear-gradient(135deg, #134e5e 0%, #71b280 100%);
  }

  &.factory-bg {
    background: linear-gradient(135deg, #2c3e50 0%, #3498db 100%);
  }

  &.driver-bg {
    background: linear-gradient(135deg, #ee9ca7 0%, #ffdde1 100%);
  }

  &.shop-bg {
    background: linear-gradient(135deg, #2193b0 0%, #6dd5ed 100%);
  }

  &.default-bg {
    background: linear-gradient(135deg, #373B44 0%, #4286f4 100%);
  }

  .user-info-card {
    margin-bottom: 25px;
    border-radius: 15px;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(31, 38, 135, 0.15);
    border: 1px solid rgba(255, 255, 255, 0.18);
    transition: transform 0.3s ease;

    &:hover {
      transform: translateY(-5px);
    }

    .user-info {
      padding: 20px;

      .info-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 20px;
        border-bottom: 2px solid #eee;
        padding-bottom: 15px;

        h2 {
          color: #2c3e50;
          margin: 0;
          font-size: 24px;
          font-weight: 600;
        }

        .role-icon {
          i {
            font-size: 32px;
            color: #16a085;
          }
        }
      }

      .info-content {
        display: flex;
        align-items: center;

        .info-item {
          display: flex;
          align-items: center;

          i {
            margin-right: 10px;
            font-size: 18px;
            color: #16a085;
          }

          .user-text,
          .role-text {
            font-size: 16px;
            color: #2c3e50;

            strong {
              font-weight: 600;
              color: #16a085;
              margin-left: 5px;
            }
          }
        }

        .el-divider {
          margin: 0 30px;
          height: 25px;
        }
      }
    }
  }

  .main-data-card {
    border-radius: 15px;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px rgba(31, 38, 135, 0.15);
    border: 1px solid rgba(255, 255, 255, 0.18);
    transition: transform 0.3s ease;

    &:hover {
      transform: translateY(-5px);
    }

    .custom-table {
      background-color: transparent;

      &::before {
        display: none;
      }

      .el-table__expanded-cell {
        background: rgba(255, 255, 255, 0.8);
      }

      .el-table__row {
        transition: all 0.3s ease;

        &:hover {
          background-color: rgba(255, 255, 255, 0.1) !important;
          transform: scale(1.01);
        }
      }
    }

    .demo-table-expand {
      padding: 20px;

      .info-section {
        margin-bottom: 30px;
        background: rgba(255, 255, 255, 0.7);
        border-radius: 10px;
        padding: 20px;
        transition: all 0.3s ease;

        &:hover {
          background: rgba(255, 255, 255, 0.9);
          transform: translateY(-2px);
        }

        .section-header {
          display: flex;
          align-items: center;
          margin-bottom: 15px;

          i {
            font-size: 24px;
            margin-right: 10px;
          }
        }

        .trace-text {
          font-size: 16px;
          font-weight: 600;
        }
      }
    }
  }

  .trace-code {
    background: rgba(64, 158, 255, 0.1);
    color: #409EFF;
    padding: 6px 12px;
    border-radius: 4px;
    font-family: monospace;
    font-weight: 600;
  }

  .el-tag {
    transition: all 0.3s ease;

    &:hover {
      transform: scale(1.05);
    }
  }
}
</style>
