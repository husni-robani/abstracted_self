<template>
  <div class="lg:max-w-4xl mx-auto">
    <div class="w-full lg:max-w-md h-5 border-b-2 border-gray-400 mb-6">
      <span
        class="bg-white pr-4 font-bold text-gray-700 font-roboto text-2xl md:text-4xl"
        >My Bio</span
      >
    </div>
    <!-- my bio content -->
    <div class="lg:flex">
      <div class="lg:mr-20 lg:w-1/2 text-lg text-justify md:text-left">
        <p>
          <br />
          Hello! i'm bani, a fresh graduate in informatics, passionate about
          turning ideas into impactful software solution
          <br /><br />
          Throughout my studies and internship program, i'v gained hands-on
          experience in developing various applications, especially websites,
          utilizing multiple programming languages and frameworks such as
          Python, JavaScript, PHP, Flask, Laravel, Nest js, Vue, and React.
          Lately, I've been diving into backend development, focusing on
          building robust APIs and implementing coding best practices to deliver
          reliable and efficient solutions. for model services
          <br /><br />
          I also enjoy of coding and excited to bring my skills and knowledge to
          the industry and contribute to the development of innovative
          applications and websites
          <br /><br />
        </p>
      </div>
      <!-- bio image -->
      <div class="mx-auto mt-6 lg:mt-0">
        <div
          class="flex relative mx-auto w-60 h-60 sm:w-80 sm:h-80 border-2 border-black rounded-md"
        >
          <img
            src="/images/profile.jpg"
            alt="profile"
            class="object object-scale-down h-60 w-60 absolute transition-transform hover:-translate-x-1 hover:-translate-y-1 -top-2 -left-[6px] sm:-top-3 sm:-left-[10px] aspect-square sm:w-80 sm:h-80 rounded-md"
          />
        </div>
      </div>
    </div>
    <!-- programming langauage -->
    <div class="flex flex-col my-10 mx-4 md:m-14 items-center">
      <!-- title -->
      <div class="text-xl text-center font-semibold">
        <h1>Technologies I've Used and Still Catch Up</h1>
        <hr class="w-1/2 mt-2 mx-auto h-2 border-t border-gray-400" />
      </div>
      <!-- programming langauage list -->
      <div class="mt-4">
        <div class="flex flex-wrap items-center justify-center">
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">Javascript</div>
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">Python</div>
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">PHP</div>
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">Laravel</div>
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">Nest js</div>
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">Flask</div>
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">Tailwind</div>
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">React js</div>
          <div class="bg-gray-200 shadow-md p-1 m-1 rounded-md">Vue js</div>
        </div>
        <div></div>
      </div>
    </div>
  </div>
</template>

<script setup></script>
