<template>
  <div id="diags">
    <section>
      <h3 class="title-text">Search Most Popular Repositories in GitHub</h3>
      <div class="todo-create">
        <section class="input">
          <div class="input">
            <v-text-field label="Fill language" v-model="inputField" v-on:keyup.enter="libResponse" hide-details="auto" > </v-text-field>
          </div>
        </section>
        <br />
        <div>
          <v-btn rounded="lg" color="primary" @click="libResponse">Run</v-btn>
          <v-dialog v-model="dialog" width="auto">
            <v-card>
              <v-card-text>
                Вы не ввели значение в строку ввода. Пожалуйста введите значение.
              </v-card-text>
              <v-card-actions>
                <v-btn color="primary" block @click="dialog = false">Close dialog</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </div>
      </div>
    </section>
  </div>
  <br />
  <div class="container">
    <v-progress-circular v-if="isLoading" size="70" width="7" color="purple" indeterminate>
    </v-progress-circular>
    <div class="container" v-else>
      <Bar :data="chartData"></Bar>
    </div>
  </div>
</template>

<script>
import { Language, GetRepos } from '@/protos/answer_pb'
import {todoServiceClient} from '@/protos/answer_grpc_web_pb'
import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'
import {object} from "google-protobuf";

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)


export default {
  name: 'Diag',
  components: { Bar },
  data () {
    return {
      chartData: {
        labels: [ 'January', 'February', 'March'],
        datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: [40, 20, 12]
          }
        ]
      },
      isLoading: false,
      dialog: false,
      inputField: "",
    }
  },
  created () {
    this.client = new todoServiceClient("http://localhost:8080", null, null)
    this.getRepos()
  },
  methods: {
    randomInteger: function (max) {
      return Math.floor(Math.random()*(max + 1));
    },
    randomRgbColor: function () {
      let r = this.randomInteger(255)
      let g = this.randomInteger(255)
      let b = this.randomInteger(255)
      return 'rgba'+'(' + r + ',' + g + ',' + b + ',' + '0.6' + ')'
    },
    getRepos: function () {
      let getRequest = new GetRepos();
      var i = 0
      var colors = []
      this.client.getRepos(getRequest, {}, (err, response) => {
        let starr = response.toObject().librariesList.map(x => {
          return x.stars
        })
        let labeler = response.toObject().librariesList.map(c => {
          return c.name
        })
        for(i=0; i < 30; i++) {
          colors.push(this.randomRgbColor())
        }

        this.chartData = {
          labels: labeler,
          datasets: [
            {
              label: "Stars",
              backgroundColor: colors,
              data: starr,
            }
          ]
        }
        this.isLoading = false
        console.log(this.chartData.datasets)
      })
    },
    libResponse: function () {
      this.isLoading = true
      let request = new Language();
      if (this.inputField === "") {
        this.dialog = true
      } else {
        request.setName(this.inputField)
        this.client.libResponse(request, {}, () => {
          this.inputField = ""
          this.getRepos()
        })
      }
    }
  },
}
</script>

<style scoped>

</style>
