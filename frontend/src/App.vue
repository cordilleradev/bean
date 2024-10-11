<template>
  <div>
    <pre>{{ formattedInfo }}</pre>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
import { APIClient } from "./services/apiClient"; // Adjust the import path as needed
import { InfoResponse, ExchangeInfo } from "./services/types";

@Options({
  components: {},
})
export default class App extends Vue {
  private apiClient!: APIClient;
  formattedInfo = ""; // Added a class property for formattedInfo

  data() {
    return {
      exchangeInfo: {} as { [key: string]: ExchangeInfo },
      loading: true,
      error: null as string | null,
    };
  }

  created() {
    this.apiClient = new APIClient("https://bean-4kc6.onrender.com");
  }
  async mounted() {
    try {
      const info: InfoResponse = await this.apiClient.getInfo();
      this.formattedInfo = JSON.stringify(info, null, 2);
      console.log(info);
    } catch (err) {
      this.formattedInfo = JSON.stringify(err, null, 2);
    }
  }
}
</script>

<style scoped>
/* Add any component-specific styles here */
</style>
