var gulp = require('gulp');
var argv = require('yargs').argv;
var jsonModify = require("gulp-json-modify");
console.log(argv.tag)
gulp.task('build', function () {
  return gulp.src(['eb-deploy/Dockerrun.aws.json'])
    .pipe(jsonModify({
      key: 'Image.Name',
      value: 'ngocnd0607/circleci-demo:'+argv.tag
    }))
    .pipe(gulp.dest('eb-deploy/'))
})
gulp.task('default', ['build']);