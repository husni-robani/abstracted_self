<template>
  <div class="lg:max-w-4xl mx-auto mt-52">
    <!-- title -->
    <div class="w-full h-5 border-b-2 border-gray-400 text-2xl mb-16">
      <span
        class="bg-white pr-4 font-bold text-gray-700 font-roboto text-2xl md:text-4xl"
        >project I've built</span
      >
    </div>
    <!-- projects list -->
    <div
      id="projects"
      class="flex flex-col items-center relative gap-y-28 lg:gap-y-40"
    >
      <!-- project 1 -->
      <div
        class="flex flex-col items-center w-full lg:items-start lg:flex-row lg:w-full"
      >
        <!-- image -->

        <div class="w-full max-w-xl">
          <div
            class="bg-cover bg-center aspect-w-16 aspect-h-9 rounded-2xl border"
            style="background-image: url('/images/LSTM_architecture.png')"
          ></div>
        </div>
        <!-- content -->
        <div
          class="w-full max-w-2xl mt-2 lg:mt-0 lg:w-1/2 lg:absolute bg-transparent right-0 lg:text-right"
        >
          <!-- title -->
          <div class="flex flex-col lg:items-end">
            <p
              class="text-sm md:text-sm font-roboto font-extralight text-gray-400"
            >
              featured project
            </p>
            <h1 class="text-lg md:text-xl font-bold lg:w-1/2">
              Text Classifier Service
            </h1>
          </div>
          <!-- description -->
          <div
            class="mt-2 lg:bg-gray-100 lg:p-2 lg:rounded-md lg:shadow-md lg:mt-4 text-base"
          >
            <p>
              Developed a deep learning model using LSTM algorithm in PyTorch
              for abstract classification, integrated with a Flask API.
              Completed as a thesis project, managing all aspects from model
              development to API deployment and documentation.
            </p>
          </div>
          <!-- stack -->
          <div
            class="flex flex-row lg:flex-row-reverse gap-2 bg-gray-100 p-2 rounded-md shadow-sm lg:bg-transparent lg:p-0 lg:shadow-none lg:rounded-none mt-2 md:mt-4"
          >
            <span
              class="flex flex-wrap flex-row gap-2 lg:w-2/3 text-sm font-roboto font-extralight text-gray-400"
            >
              <span>Python</span>
              <span>Flask</span>
              <span>Pytorch</span>
            </span>
          </div>
          <!-- link -->
          <div class="flex lg:flex-row-reverse gap-2 mt-2 lg:mt-4">
            <a
              href="https://github.com/husni-robani/service_api_for_lstm_model"
              target="_blank"
              ><svg
                class="fill-gray-800 hover:fill-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 30 30"
                width="30px"
                height="30px"
              >
                <path
                  d="M15,3C8.373,3,3,8.373,3,15c0,5.623,3.872,10.328,9.092,11.63C12.036,26.468,12,26.28,12,26.047v-2.051 c-0.487,0-1.303,0-1.508,0c-0.821,0-1.551-0.353-1.905-1.009c-0.393-0.729-0.461-1.844-1.435-2.526 c-0.289-0.227-0.069-0.486,0.264-0.451c0.615,0.174,1.125,0.596,1.605,1.222c0.478,0.627,0.703,0.769,1.596,0.769 c0.433,0,1.081-0.025,1.691-0.121c0.328-0.833,0.895-1.6,1.588-1.962c-3.996-0.411-5.903-2.399-5.903-5.098 c0-1.162,0.495-2.286,1.336-3.233C9.053,10.647,8.706,8.73,9.435,8c1.798,0,2.885,1.166,3.146,1.481C13.477,9.174,14.461,9,15.495,9 c1.036,0,2.024,0.174,2.922,0.483C18.675,9.17,19.763,8,21.565,8c0.732,0.731,0.381,2.656,0.102,3.594 c0.836,0.945,1.328,2.066,1.328,3.226c0,2.697-1.904,4.684-5.894,5.097C18.199,20.49,19,22.1,19,23.313v2.734 c0,0.104-0.023,0.179-0.035,0.268C23.641,24.676,27,20.236,27,15C27,8.373,21.627,3,15,3z"
                /></svg
            ></a>
            <a
              href="https://github.com/husni-robani/LSTM-multi-class-text-classification"
              target="_blank"
              ><svg
                class="fill-gray-800 hover:fill-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 30 30"
                width="30px"
                height="30px"
              >
                <path
                  d="M15,3C8.373,3,3,8.373,3,15c0,5.623,3.872,10.328,9.092,11.63C12.036,26.468,12,26.28,12,26.047v-2.051 c-0.487,0-1.303,0-1.508,0c-0.821,0-1.551-0.353-1.905-1.009c-0.393-0.729-0.461-1.844-1.435-2.526 c-0.289-0.227-0.069-0.486,0.264-0.451c0.615,0.174,1.125,0.596,1.605,1.222c0.478,0.627,0.703,0.769,1.596,0.769 c0.433,0,1.081-0.025,1.691-0.121c0.328-0.833,0.895-1.6,1.588-1.962c-3.996-0.411-5.903-2.399-5.903-5.098 c0-1.162,0.495-2.286,1.336-3.233C9.053,10.647,8.706,8.73,9.435,8c1.798,0,2.885,1.166,3.146,1.481C13.477,9.174,14.461,9,15.495,9 c1.036,0,2.024,0.174,2.922,0.483C18.675,9.17,19.763,8,21.565,8c0.732,0.731,0.381,2.656,0.102,3.594 c0.836,0.945,1.328,2.066,1.328,3.226c0,2.697-1.904,4.684-5.894,5.097C18.199,20.49,19,22.1,19,23.313v2.734 c0,0.104-0.023,0.179-0.035,0.268C23.641,24.676,27,20.236,27,15C27,8.373,21.627,3,15,3z"
                /></svg
            ></a>
          </div>
        </div>
      </div>
      <!-- project 2 -->
      <div
        class="flex flex-col items-center w-full lg:items-start lg:flex-row-reverse lg:w-full"
      >
        <!-- image -->
        <div class="w-full max-w-xl">
          <div
            class="bg-cover bg-center aspect-w-16 aspect-h-9 rounded-2xl border"
            style="background-image: url('/images/himatif_vote.png')"
          ></div>
        </div>
        <!-- content -->
        <div
          class="w-full max-w-2xl mt-2 lg:mt-0 lg:w-1/2 lg:absolute bg-transparent left-0 lg:text-left"
        >
          <!-- title -->
          <div class="flex flex-col">
            <p
              class="text-sm md:text-sm font-roboto font-extralight text-gray-400"
            >
              featured project
            </p>
            <h1 class="text-lg md:text-xl font-bold lg:w-1/2">Vote App</h1>
          </div>
          <!-- description -->
          <div
            class="mt-2 lg:bg-gray-100 lg:p-2 lg:rounded-md lg:shadow-md lg:mt-4 text-base"
          >
            <p>
              The HIMATIF Voting System is a web application designed for
              electing organization leaders. As a full-stack engineer, I
              developed this project independently, handling all aspects from
              UI/UX design to frontend and backend development, and deployment.
            </p>
          </div>
          <!-- stack -->
          <div
            class="flex flex-row lg:flex-row gap-2 bg-gray-100 p-2 rounded-md shadow-sm lg:bg-transparent lg:p-0 lg:shadow-none lg:rounded-none mt-2 md:mt-4"
          >
            <span
              class="flex flex-wrap flex-row gap-2 lg:w-2/3 text-sm font-roboto font-extralight text-gray-400"
            >
              <span>PHP</span>
              <span>Javascript</span>
              <span>Laravel</span>
              <span>Vue Js</span>
              <span>Tailwind</span>
              <span>Inertia</span>
              <span>MySQL</span>
            </span>
          </div>
          <!-- link -->
          <div class="flex lg:flex-row gap-2 mt-2 lg:mt-4">
            <a
              href="https://github.com/husni-robani/vote-system"
              target="_blank"
              ><svg
                class="fill-gray-800 hover:fill-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 30 30"
                width="30px"
                height="30px"
              >
                <path
                  d="M15,3C8.373,3,3,8.373,3,15c0,5.623,3.872,10.328,9.092,11.63C12.036,26.468,12,26.28,12,26.047v-2.051 c-0.487,0-1.303,0-1.508,0c-0.821,0-1.551-0.353-1.905-1.009c-0.393-0.729-0.461-1.844-1.435-2.526 c-0.289-0.227-0.069-0.486,0.264-0.451c0.615,0.174,1.125,0.596,1.605,1.222c0.478,0.627,0.703,0.769,1.596,0.769 c0.433,0,1.081-0.025,1.691-0.121c0.328-0.833,0.895-1.6,1.588-1.962c-3.996-0.411-5.903-2.399-5.903-5.098 c0-1.162,0.495-2.286,1.336-3.233C9.053,10.647,8.706,8.73,9.435,8c1.798,0,2.885,1.166,3.146,1.481C13.477,9.174,14.461,9,15.495,9 c1.036,0,2.024,0.174,2.922,0.483C18.675,9.17,19.763,8,21.565,8c0.732,0.731,0.381,2.656,0.102,3.594 c0.836,0.945,1.328,2.066,1.328,3.226c0,2.697-1.904,4.684-5.894,5.097C18.199,20.49,19,22.1,19,23.313v2.734 c0,0.104-0.023,0.179-0.035,0.268C23.641,24.676,27,20.236,27,15C27,8.373,21.627,3,15,3z"
                /></svg
            ></a>
            <a href="https://project1.husnirb.tech" target="_blank"
              ><svg
                class="fill-gray-800 hover:fill-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 64 64"
                width="30px"
                height="30px"
              >
                <path
                  d="M 40 10 C 38.896 10 38 10.896 38 12 C 38 13.104 38.896 14 40 14 L 47.171875 14 L 30.585938 30.585938 C 29.804938 31.366938 29.804938 32.633063 30.585938 33.414062 C 30.976938 33.805063 31.488 34 32 34 C 32.512 34 33.023063 33.805062 33.414062 33.414062 L 50 16.828125 L 50 24 C 50 25.104 50.896 26 52 26 C 53.104 26 54 25.104 54 24 L 54 12 C 54 10.896 53.104 10 52 10 L 40 10 z M 18 12 C 14.691 12 12 14.691 12 18 L 12 46 C 12 49.309 14.691 52 18 52 L 46 52 C 49.309 52 52 49.309 52 46 L 52 34 C 52 32.896 51.104 32 50 32 C 48.896 32 48 32.896 48 34 L 48 46 C 48 47.103 47.103 48 46 48 L 18 48 C 16.897 48 16 47.103 16 46 L 16 18 C 16 16.897 16.897 16 18 16 L 30 16 C 31.104 16 32 15.104 32 14 C 32 12.896 31.104 12 30 12 L 18 12 z"
                />
              </svg>
            </a>
          </div>
        </div>
      </div>
      <!-- project 3 -->
      <div
        class="flex flex-col items-center w-full lg:items-start lg:flex-row lg:w-full"
      >
        <!-- image -->

        <div class="w-full max-w-xl">
          <div
            class="bg-cover bg-center aspect-w-16 aspect-h-9 rounded-2xl border"
            style="background-image: url('/images/image.png')"
          ></div>
        </div>
        <!-- content -->
        <div
          class="w-full max-w-2xl mt-2 lg:mt-0 lg:w-1/2 lg:absolute bg-transparent right-0 lg:text-right"
        >
          <!-- title -->
          <div class="flex flex-col lg:items-end">
            <p
              class="text-sm md:text-sm font-roboto font-extralight text-gray-400"
            >
              featured project
            </p>
            <h1 class="text-lg md:text-xl font-bold lg:w-1/2">
              API Services - Zakat APP
            </h1>
          </div>
          <!-- description -->
          <div
            class="mt-2 lg:bg-gray-100 lg:p-2 lg:rounded-md lg:shadow-md lg:mt-4 text-base"
          >
            <p>
              This project was developed for the PKM program. Riyadhul Jannah
              Zakat App, where I served as the backend engineer, focused on
              developing the backend system supporting the application
            </p>
          </div>
          <!-- stack -->
          <div
            class="flex flex-row lg:flex-row-reverse gap-2 bg-gray-100 p-2 rounded-md shadow-sm lg:bg-transparent lg:p-0 lg:shadow-none lg:rounded-none mt-2 md:mt-4"
          >
            <span
              class="flex flex-wrap flex-row gap-2 lg:w-2/3 text-sm font-roboto font-extralight text-gray-400"
            >
              <span>PHP</span>
              <span>Laravel</span>
              <span>MySQL</span>
            </span>
          </div>
          <!-- link -->
          <div class="flex lg:flex-row-reverse gap-2 mt-2 lg:mt-4">
            <a
              href="https://github.com/husni-robani/riyadhul-zakat-service"
              target="_blank"
              ><svg
                class="fill-gray-800 hover:fill-gray-500"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 30 30"
                width="30px"
                height="30px"
              >
                <path
                  d="M15,3C8.373,3,3,8.373,3,15c0,5.623,3.872,10.328,9.092,11.63C12.036,26.468,12,26.28,12,26.047v-2.051 c-0.487,0-1.303,0-1.508,0c-0.821,0-1.551-0.353-1.905-1.009c-0.393-0.729-0.461-1.844-1.435-2.526 c-0.289-0.227-0.069-0.486,0.264-0.451c0.615,0.174,1.125,0.596,1.605,1.222c0.478,0.627,0.703,0.769,1.596,0.769 c0.433,0,1.081-0.025,1.691-0.121c0.328-0.833,0.895-1.6,1.588-1.962c-3.996-0.411-5.903-2.399-5.903-5.098 c0-1.162,0.495-2.286,1.336-3.233C9.053,10.647,8.706,8.73,9.435,8c1.798,0,2.885,1.166,3.146,1.481C13.477,9.174,14.461,9,15.495,9 c1.036,0,2.024,0.174,2.922,0.483C18.675,9.17,19.763,8,21.565,8c0.732,0.731,0.381,2.656,0.102,3.594 c0.836,0.945,1.328,2.066,1.328,3.226c0,2.697-1.904,4.684-5.894,5.097C18.199,20.49,19,22.1,19,23.313v2.734 c0,0.104-0.023,0.179-0.035,0.268C23.641,24.676,27,20.236,27,15C27,8.373,21.627,3,15,3z"
                /></svg
            ></a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup></script>
