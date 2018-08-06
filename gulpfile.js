var gulp = require('gulp');
var argv = require('yargs').argv;
var jsonModify = require("gulp-json-modify");
gulp.task('build', function () {
  return gulp.src(['eb-deploy/Dockerrun.aws.json'])
    .pipe(jsonModify({
      key: 'Image.Name',
      value: argv.url
    }))
    .pipe(gulp.dest('eb-deploy/'))
})
gulp.task('default', ['build']);