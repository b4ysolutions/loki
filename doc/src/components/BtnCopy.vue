<template>
  <button
    type="button"
    class="copy-btn"
    v-clipboard:copy="text"
    v-clipboard:success="onCopy"
    v-clipboard:error="onError"
  >
    <content-copy-icon />
  </button>
</template>

<script>
import ContentCopyIcon from "vue-material-design-icons/ContentCopy.vue";

export default {
  name: "BtnCopy",
  props: {
    text: String,
  },
  components: {
    ContentCopyIcon,
  },
  data() {
    return {};
  },
  methods: {
    onCopy: function() {
      this.$notify({
        group: 'loki',
        type: 'success',
        title: 'Successfully Copied!'
      })
    },
    onError: function() {
      this.$notify({
        group: 'loki',
        type: 'error',
        title: 'Failed to Copy!'
      })
    },
    parseToObject: function(str) {
      return JSON.parse(str);
    },
  },
};

</script>

<style scoped></style>
