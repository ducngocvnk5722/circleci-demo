var gulp = require('gulp');

gulp.task('build', function(){
  return gulp.src('client/templates/*.pug')
    .pipe(pug())
    .pipe(gulp.dest('build/html'))
});
gulp.task('default', ['build']);