<script lang="ts">
  import { serverUrl } from './Stores';

  let value: string;

  function onSubmit() {
    serverUrl.set(value);
  }
</script>

<div
  class="flex place-items-center justify-center text-center h-screen flex-col"
>
  <h1 class="text-4xl mb-20 block">Select Server</h1>
  <form on:submit|preventDefault={onSubmit}>
    <input
      class="border border-black rounded px-3 py-2"
      type="text"
      placeholder="Enter server address"
      bind:value
    />
    <input
      class="px-3 py-2 bg-neutral-300 rounded ml-5"
      type="submit"
      value="Connect"
    />
  </form>
</div>
