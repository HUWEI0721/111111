<!-- Start of Selection -->
<template>
  <div class="uplink-container" :class="roleClass">
    <el-card class="user-info-card" shadow="hover">
      <div class="user-info">
        <div class="info-header">
          <h2>用户信息</h2>
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

    <el-card class="main-form-card" shadow="hover">
      <el-form ref="form" :model="tracedata" label-width="120px" size="small">

        <el-form-item v-show="userType != '茅台厂商' & userType != '消费者'" label="溯源码:" class="trace-code-item">
          <el-input v-model="tracedata.traceability_code" placeholder="请输入溯源码">
            <template slot="prepend"><i class="el-icon-key"></i></template>
          </el-input>
        </el-form-item>

        <!-- 茅台厂商信息录入区域 -->
        <div v-show="userType == '茅台厂商'" class="role-form-section">
          <h3><i class="el-icon-ship"></i> 茅台厂商信息录入 <i class="el-icon-fishing role-specific-icon"></i></h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="茅台酒名称:">
                <el-input v-model="tracedata.Brewer_input.Moutai_name" placeholder="请输入茅台酒名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="产地:">
                <el-input v-model="tracedata.Brewer_input.Moutai_origin" placeholder="请输入产地" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="酿造开始时间:">
                <el-date-picker v-model="tracedata.Brewer_input.Moutai_start" type="datetime" placeholder="选择日期时间"
                  style="width: 100%" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="酿造完成时间:">
                <el-date-picker v-model="tracedata.Brewer_input.Moutai_end" type="datetime" placeholder="选择日期时间"
                  style="width: 100%" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="酿酒师名称:">
                <el-input v-model="tracedata.Brewer_input.Brewer_name" placeholder="请输入酿酒师名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="制造时长:">
                <el-input :value="calculateDuration" readonly>
                  <template slot="append">小时</template>
                </el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- 酿造厂信息录入区域 -->
        <div v-show="userType == '酿造厂'" class="role-form-section">
          <h3><i class="el-icon-office-building"></i> 酿造厂信息录入 <i class="el-icon-box role-specific-icon"></i></h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="商品名称:">
                <el-input v-model="tracedata.Distillery_input.Moutai_productName" placeholder="请输入商品名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="酿造批次:">
                <el-input v-model="tracedata.Distillery_input.Moutai_batch" placeholder="请输入酿造批次" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="酿造时间:">
                <el-date-picker v-model="tracedata.Distillery_input.Moutai_productionTime" type="datetime"
                  placeholder="选择日期时间" style="width: 100%" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="酿造厂名称:">
                <el-input v-model="tracedata.Distillery_input.Distillery_name" placeholder="请输入酿造厂名称" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="联系电话:">
                <el-input v-model="tracedata.Distillery_input.Distillery_phone" placeholder="请输入联系电话" />
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- 物流司机信息录入区域 -->
        <div v-show="userType == '物流司机信息'" class="role-form-section">
          <h3><i class="el-icon-truck"></i> 物流信息录入 <i class="el-icon-van role-specific-icon"></i></h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="司机姓名:">
                <el-input v-model="tracedata.Driver_input.Driver_name" placeholder="请输入司机姓名" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="车厢温度:">
                <el-input v-model="tracedata.Driver_input.Driver_age" placeholder="请输入车厢温度">
                  <template slot="append">°C</template>
                </el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="联系电话:">
                <el-input v-model="tracedata.Driver_input.Driver_phone" placeholder="请输入联系电话" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="车牌号:">
                <el-input v-model="tracedata.Driver_input.Driver_carNumber" placeholder="请输入车牌号" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="运输工具:">
                <el-input v-model="tracedata.Driver_input.Driver_transport" placeholder="请输入运输工具" />
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- 销售点信息录入区域 -->
        <div v-show="userType == '销售终端'" class="role-form-section">
          <h3><i class="el-icon-shopping-cart-full"></i> 销售点信息录入 <i class="el-icon-goods role-specific-icon"></i></h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="存入时间:">
                <el-date-picker v-model="tracedata.Shop_input.Shop_storeTime" type="datetime" placeholder="选择日期时间"
                  style="width: 100%" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="销售时间:">
                <el-date-picker v-model="tracedata.Shop_input.Shop_sellTime" type="datetime" placeholder="选择日期时间"
                  style="width: 100%" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="销售点名称:">
                <el-input v-model="tracedata.Shop_input.Shop_shopName" placeholder="请输入销售点名称" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="销售点位置:">
                <el-input v-model="tracedata.Shop_input.Shop_shopAddress" placeholder="请输入销售点位置" />
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="销售点电话:">
                <el-input v-model="tracedata.Shop_input.Shop_shopPhone" placeholder="请输入销售点电话" />
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- 图片上传区域 -->
        <div v-show="userType != '消费者'" class="upload-section">
          <h3><i class="el-icon-picture"></i> 图片上传</h3>
          <el-upload class="upload-area" action="#" :on-preview="handlePreview" :on-remove="handleRemove"
            :before-remove="beforeRemove" :on-change="handleChange" :file-list="fileList" :auto-upload="false"
            :limit="1" list-type="picture-card">
            <template v-if="fileList.length < 1">
              <i class="el-icon-plus upload-icon"></i>
            </template>
          </el-upload>
          <div class="upload-tip" v-if="fileList.length < 1">
            <i class="el-icon-info-circle"></i>
            <span>图片大小不超过 500KB，建议尺寸 1080x1080px</span>
          </div>
        </div>

        <div class="form-footer">
          <el-alert v-show="userType == '消费者'" title="消费者没有权限录入！请使用溯源功能!" type="warning" :closable="false">
          </el-alert>
          <el-button v-show="userType != '消费者'" type="success" icon="el-icon-upload2" @click="submittracedata">
            提交数据上链
          </el-button>
        </div>
      </el-form>
    </el-card>

    <!-- 图片预览对话框 -->
    <el-dialog :visible.sync="previewVisible">
      <img width="100%" :src="previewUrl" alt="Preview">
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { uplink } from '@/api/trace'
import Compressor from 'compressorjs';

export default {
  name: 'Uplink',
  data() {
    return {
      tracedata: {
        traceability_code: '',
        Brewer_input: {
          Moutai_name: '',
          Moutai_origin: '',
          Moutai_start: '',
          Moutai_end: '',
          Brewer_name: ''
        },
        Distillery_input: {
          Moutai_productName: '',
          Moutai_batch: '',
          Moutai_productionTime: '',
          Distillery_name: '',
          Distillery_phone: ''
        },
        Driver_input: {
          Driver_name: '',
          Driver_age: '',
          Driver_phone: '',
          Driver_carNumber: '',
          Driver_transport: ''
        },
        Shop_input: {
          Shop_storeTime: '',
          Shop_sellTime: '',
          Shop_shopName: '',
          Shop_shopAddress: '',
          Shop_shopPhone: ''
        }
      },
      fileList: [],
      loading: false,
      previewVisible: false,
      previewUrl: ''
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ]),
    calculateDuration() {
      if (this.tracedata.Brewer_input.Moutai_start && this.tracedata.Brewer_input.Moutai_end) {
        const startTime = new Date(this.tracedata.Brewer_input.Moutai_start).getTime()
        const endTime = new Date(this.tracedata.Brewer_input.Moutai_end).getTime()
        const diffHours = Math.round((endTime - startTime) / (1000 * 60 * 60))
        return diffHours >= 0 ? diffHours : 0
      }
      return 0
    },
    roleClass() {
      switch (this.userType) {
        case '茅台厂商':
          return 'fisherman-bg'
        case '酿造厂':
          return 'factory-bg'
        case '物流司机信息':
          return 'driver-bg'
        case '销售终端':
          return 'shop-bg'
        default:
          return 'default-bg'
      }
    },
    roleIcon() {
      switch (this.userType) {
        case '茅台厂商':
          return 'el-icon-ship'
        case '酿造厂':
          return 'el-icon-office-building'
        case '物流司机信息':
          return 'el-icon-truck'
        case '销售终端':
          return 'el-icon-shopping-cart-full'
        default:
          return 'el-icon-user'
      }
    }
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      this.previewUrl = file.url;
      this.previewVisible = true;
    },
    handleChange(file, fileList) {
      new Compressor(file.raw, {
        quality: 0.2, // 压缩质量，0到1之间
        convertSize: Infinity, // 确保所有图片都被转换
        mimeType: 'image/jpeg', // 转换为jpg格式

        success: (compressedFile) => {
          this.fileList = [{ raw: compressedFile, url: URL.createObjectURL(compressedFile) }];
        },
        error(err) {
          console.error(err.message);
        },
      });
    },
    beforeRemove(file, fileList) {
      return this.$confirm(`确定移除 ${file.name}？`);
    },
    submittracedata() {
      console.log(this.tracedata)
      const loading = this.$loading({
        lock: true,
        text: '数据上链中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      var formData = new FormData()
      formData.append('traceability_code', this.tracedata.traceability_code)
      // 根据不同的用户给arg1、arg2、arg3..赋值,
      switch (this.userType) {
        case '茅台厂商':
          formData.append('arg1', this.tracedata.Brewer_input.Moutai_name)
          formData.append('arg2', this.tracedata.Brewer_input.Moutai_origin)
          formData.append('arg3', this.tracedata.Brewer_input.Moutai_start)
          formData.append('arg4', this.tracedata.Brewer_input.Moutai_end)
          formData.append('arg5', this.tracedata.Brewer_input.Brewer_name)
          //formData.append('image', this.tracedata.Brewer_input.Sf_duration)
          break
        case '酿造厂':
          formData.append('arg1', this.tracedata.Distillery_input.Moutai_productName)
          formData.append('arg2', this.tracedata.Distillery_input.Moutai_batch)
          formData.append('arg3', this.tracedata.Distillery_input.Moutai_productionTime)
          formData.append('arg4', this.tracedata.Distillery_input.Distillery_name)
          formData.append('arg5', this.tracedata.Distillery_input.Distillery_phone)
          break
        case '物流司机信息':
          formData.append('arg1', this.tracedata.Driver_input.Driver_name)
          formData.append('arg2', this.tracedata.Driver_input.Driver_age)
          formData.append('arg3', this.tracedata.Driver_input.Driver_phone)
          formData.append('arg4', this.tracedata.Driver_input.Driver_carNumber)
          formData.append('arg5', this.tracedata.Driver_input.Driver_transport)
          break
        case '销售终端':
          formData.append('arg1', this.tracedata.Shop_input.Shop_storeTime)
          formData.append('arg2', this.tracedata.Shop_input.Shop_sellTime)
          formData.append('arg3', this.tracedata.Shop_input.Shop_shopName)
          formData.append('arg4', this.tracedata.Shop_input.Shop_shopAddress)
          formData.append('arg5', this.tracedata.Shop_input.Shop_shopPhone)
          break
      }
      if (this.fileList.length > 0) {
        const reader = new FileReader();
        reader.readAsDataURL(this.fileList[0].raw);
        reader.onload = () => {
          formData.append('image', reader.result);
          console.log(formData)
          uplink(formData).then(res => {
            if (res.code === 200) {
              loading.close()
              this.$message({
                message: '上链成功，交易ID：' + res.txid + '\n溯源码：' + res.traceability_code,
                type: 'success'
              })
              this.resetForm()
            } else {
              loading.close()
              this.$message({
                message: '上链失败',
                type: 'error'
              })
            }
          }).catch(err => {
            loading.close()
            console.log(err)
          })
        };
      } else {
        uplink(formData).then(res => {
          if (res.code === 200) {
            loading.close()
            this.$message({
              message: '上链成功，交易ID：' + res.txid + '\n溯源码：' + res.traceability_code,
              type: 'success'
            })
            this.resetForm()
          } else {
            loading.close()
            this.$message({
              message: '上链失败',
              type: 'error'
            })
          }
        }).catch(err => {
          loading.close()
          console.log(err)
        })
      }
    },
    resetForm() {
      this.tracedata = {
        traceability_code: '',
        Brewer_input: {
          Moutai_name: '',
          Moutai_origin: '',
          Moutai_start: '',
          Moutai_end: '',
          Brewer_name: ''
        },
        Distillery_input: {
          Moutai_productName: '',
          Moutai_batch: '',
          Moutai_productionTime: '',
          Distillery_name: '',
          Distillery_phone: ''
        },
        Driver_input: {
          Driver_name: '',
          Driver_age: '',
          Driver_phone: '',
          Driver_carNumber: '',
          Driver_transport: ''
        },
        Shop_input: {
          Shop_storeTime: '',
          Shop_sellTime: '',
          Shop_shopName: '',
          Shop_shopAddress: '',
          Shop_shopPhone: ''
        }
      }
      this.fileList = []
    }
  }
}
</script>

<style lang="scss" scoped>
.uplink-container {
  padding: 30px;
  min-height: 100vh;

  &.fisherman-bg {
    background: linear-gradient(135deg, #1e3c72, #2a5298);
  }

  &.factory-bg {
    background: linear-gradient(135deg, #434343, #000000);
  }

  &.driver-bg {
    background: linear-gradient(135deg, #2c3e50, #3498db);
  }

  &.shop-bg {
    background: linear-gradient(135deg, #134e5e, #71b280);
  }

  &.default-bg {
    background: linear-gradient(135deg, #1a2a6c, #b21f1f, #fdbb2d);
  }

  .user-info-card {
    margin-bottom: 25px;
    border-radius: 15px;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);

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
            letter-spacing: 0.5px;

            strong {
              font-weight: 600;
              color: #16a085;
              margin-left: 5px;
              font-size: 17px;
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

  .main-form-card {
    border-radius: 15px;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.1);

    .role-form-section {
      background: rgba(236, 240, 241, 0.5);
      border-radius: 12px;
      padding: 25px;
      margin-bottom: 25px;
      border: 1px solid rgba(189, 195, 199, 0.3);
      transition: all 0.3s ease;

      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
      }

      h3 {
        color: #2c3e50;
        margin: 0 0 25px 0;
        font-size: 20px;
        font-weight: 600;
        display: flex;
        align-items: center;
        justify-content: space-between;

        i {
          color: #16a085;
          font-size: 24px;

          &.role-specific-icon {
            font-size: 28px;
            color: #2980b9;
          }
        }
      }
    }

    .upload-section {
      background: rgba(236, 240, 241, 0.5);
      border-radius: 12px;
      padding: 25px;
      margin-bottom: 25px;
      border: 1px solid rgba(189, 195, 199, 0.3);
      text-align: center;
      transition: all 0.3s ease;

      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
      }

      h3 {
        color: #2c3e50;
        margin: 0 0 25px 0;
        font-size: 20px;
        font-weight: 600;
        display: flex;
        align-items: center;
        justify-content: center;

        i {
          margin-right: 10px;
          color: #16a085;
          font-size: 24px;
        }
      }

      .upload-area {
        display: flex;
        justify-content: center;
        margin-bottom: 20px;

        ::v-deep .el-upload--picture-card {
          width: 180px;
          height: 180px;
          line-height: 180px;
          border-radius: 12px;
          border: none;
          background: transparent;
          transition: all 0.3s ease;

          &:hover {
            border-color: #409EFF;
            color: #409EFF;
            transform: scale(1.02);
          }
        }

        .upload-icon {
          font-size: 32px;
          color: #8c939d;
          margin-bottom: 10px;
        }

        .upload-text {
          display: flex;
          flex-direction: column;
          align-items: center;
          color: #606266;
          font-size: 14px;

          .upload-hint {
            margin-top: 8px;
            font-size: 12px;
            color: #909399;
          }
        }
      }

      .upload-tip {
        display: flex;
        align-items: center;
        justify-content: center;
        color: #909399;
        font-size: 13px;

        i {
          margin-right: 8px;
          font-size: 16px;
          color: #E6A23C;
        }
      }

      ::v-deep .el-upload-list--picture-card {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 15px;

        .el-upload-list__item {
          width: 160px;
          height: 160px;
          border-radius: 8px;
          transition: all 0.3s ease;

          &:hover {
            transform: translateY(-3px);
          }
        }
      }
    }

    .trace-code-item {
      max-width: 500px;
    }

    .el-input {
      .el-input__inner {
        border-radius: 8px;
        border: 1px solid #dcdfe6;
        transition: all 0.3s ease;

        &:focus {
          border-color: #16a085;
          box-shadow: 0 0 0 2px rgba(22, 160, 133, 0.2);
        }
      }
    }

    .form-footer {
      margin-top: 40px;
      text-align: center;

      .el-button {
        padding: 12px 35px;
        font-size: 15px;
        border-radius: 8px;
        background: #16a085;
        border-color: #16a085;
        transition: all 0.3s ease;

        &:hover {
          background: #1abc9c;
          border-color: #1abc9c;
          transform: translateY(-2px);
          box-shadow: 0 5px 15px rgba(22, 160, 133, 0.3);
        }
      }

      .el-alert {
        margin-bottom: 25px;
        border-radius: 8px;
      }
    }
  }
}
</style>
