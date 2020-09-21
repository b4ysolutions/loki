<template>
  <b-container>
    <h4 :id="id">{{ title }}</h4>
    <p>{{ description }}</p>

    <h3>
      URL:
      <BtnCopy :text="url" />
    </h3>
    <div class="code">{{ url }}</div>

    <div
      v-bind:item="example"
      v-for="example in examples"
      v-bind:key="example.id"
    >
      <h5>{{ example.title }}</h5>
      <b-row class="examples-row">
        <b-col class="example-col">
          Request Url:
          <BtnCopy :text="example.url" />
          <div class="code">{{ example.url }}</div>
          Request Method:
          <BtnCopy :text="example.method" />
          <div class="code">{{ example.method }}</div>
          Request Body:
          <BtnCopy :text="JSON.stringify(example.body)" />
          <div class="code pre-json">
            <vue-json-pretty :path="'res'" :data="example.body">
            </vue-json-pretty>
          </div>
        </b-col>
        <b-col class="example-col">
          Response:
          <BtnCopy :text="JSON.stringify(example.response)" />
          <div class="code pre-json">
            <vue-json-pretty :path="'res'" :data="example.response">
            </vue-json-pretty>
          </div>
        </b-col>
      </b-row>
    </div>
  </b-container>
</template>

<script>
import VueJsonPretty from "vue-json-pretty";
import BtnCopy from "./BtnCopy.vue";

export default {
  name: "Endpoint",
  props: {
    id: String,
    title: String,
    description: String,
    url: String,
    examples: Array,
  },
  components: {
    VueJsonPretty,
    BtnCopy,
  },
  data() {
    return {};
  },
};
</script>

<style scoped></style>
